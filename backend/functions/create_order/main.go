package main

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Customer struct {
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
	Email     string `bson:"email" json:"email"`
	Phone     string `bson:"phone" json:"phone"`
	City      string `bson:"city" json:"city"`
	Address   string `bson:"address" json:"address"`
}

type PayloadItem struct {
	ID  string `json:"id"`
	Qty int    `json:"qty"`
}

type Payload struct {
	Customer       Customer      `json:"customer"`
	Items          []PayloadItem `json:"items"`
	PaymentMethod  string        `json:"paymentMethod"`
	ShippingMethod string        `json:"shippingMethod"`
	// NOTA: shippingCost del frontend se IGNORA — siempre se calcula en servidor
}

// Usamos interface{} para precio ya que a veces entra como string o como numero en MongoDB
type Product struct {
	ID    string      `bson:"id"`
	Name  string      `bson:"name"`
	Price interface{} `bson:"price"`
	Stock int         `bson:"stock"`
}

type Order struct {
	ID             string    `bson:"id" json:"id"`
	Customer       Customer  `bson:"customer" json:"customer"`
	Items          []bson.M  `bson:"items" json:"items"`
	Subtotal       float64   `bson:"subtotal" json:"subtotal"`
	SubtotalFormat string    `bson:"subtotal_format" json:"subtotal_format"`
	ShippingCost   float64   `bson:"shippingCost" json:"shippingCost"`
	ShippingFormat string    `bson:"shipping_format" json:"shipping_format"`
	ShippingMethod string    `bson:"shippingMethod" json:"shippingMethod"`
	Total          float64   `bson:"total" json:"total"`
	TotalFormat    string    `bson:"total_format" json:"total_format"`
	PaymentMethod  string    `bson:"paymentMethod" json:"paymentMethod"`
	Status         string    `bson:"status" json:"status"`
	ClientIP       string    `bson:"clientIp,omitempty" json:"clientIp,omitempty"`
	CreatedAt      time.Time `bson:"createdAt" json:"createdAt"`
	// Campos para integración Wompi (se llenan después)
	WompiTransactionID string `bson:"wompiTransactionId,omitempty" json:"wompiTransactionId,omitempty"`
	WompiStatus        string `bson:"wompiStatus,omitempty" json:"wompiStatus,omitempty"`
}

// Configura CORS
func corsHeaders() map[string]string {
	return map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Headers": "Content-Type",
		"Access-Control-Allow-Methods": "POST, OPTIONS",
		"Content-Type":                 "application/json",
	}
}

// Parsear el string con formato "$15.000 COP" a número
// Dependiendo de cómo lo tengan en DB, pero vamos a asumir que podría estar limpio o sucio
func parsePrice(v interface{}) float64 {
	switch val := v.(type) {
	case float64:
		return val
	case int32:
		return float64(val)
	case int64:
		return float64(val)
	case string:
		// Remover signos $ , puntos y espacios
		s := strings.ReplaceAll(val, "$", "")
		s = strings.ReplaceAll(s, ".", "")
		s = strings.ReplaceAll(s, " COP", "")
		s = strings.TrimSpace(s)
		var num float64
		fmt.Sscanf(s, "%f", &num)
		return num
	default:
		return 0
	}
}

func formatNumberIntl(num float64) string {
	s := fmt.Sprintf("%.0f", num)
	n := len(s)
	if n <= 3 {
		return s
	}
	var res []string
	for i := n; i > 0; i -= 3 {
		if i-3 >= 0 {
			res = append([]string{s[i-3 : i]}, res...)
		} else {
			res = append([]string{s[0:i]}, res...)
		}
	}
	return strings.Join(res, ".")
}

func getClientIP(req events.APIGatewayProxyRequest) string {
	for k, v := range req.Headers {
		lower := strings.ToLower(k)
		if lower == "x-nf-client-connection-ip" && v != "" {
			return v
		}
	}
	for k, v := range req.Headers {
		lower := strings.ToLower(k)
		if lower == "x-forwarded-for" && v != "" {
			parts := strings.Split(v, ",")
			return strings.TrimSpace(parts[0])
		}
	}
	if req.RequestContext.Identity.SourceIP != "" {
		return req.RequestContext.Identity.SourceIP
	}
	return "desconocida"
}

// calculateShippingCost determina el costo de envío SIEMPRE en el servidor
// NUNCA confía en valores enviados desde el frontend
func calculateShippingCost(shippingMethod string) float64 {
	switch shippingMethod {
	case "express_valle", "valle":
		return 10000
	case "express_alrededores", "alrededores":
		return 15000
	case "express_nacional", "nacional":
		return 20000
	default:
		// Default seguro: zona más económica
		return 10000
	}
}

// generateOrderID genera un ID de orden con crypto/rand en vez de math/rand
func generateOrderID() string {
	n, err := rand.Int(rand.Reader, big.NewInt(90000))
	if err != nil {
		// Fallback ultra-seguro usando timestamp
		return fmt.Sprintf("ORD-%d", time.Now().UnixNano()%90000+10000)
	}
	return fmt.Sprintf("ORD-%d", n.Int64()+10000)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Preflight options request (CORS)
	if request.HTTPMethod == "OPTIONS" {
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Headers:    corsHeaders(),
			Body:       "OK",
		}, nil
	}

	// 1. Parsear el Payload
	var body Payload
	if err := json.Unmarshal([]byte(request.Body), &body); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400, Headers: corsHeaders(), Body: `{"error": "JSON Inválido"}`}, nil
	}

	if len(body.Items) == 0 {
		return events.APIGatewayProxyResponse{StatusCode: 400, Headers: corsHeaders(), Body: `{"error": "El carrito está vacío."}`}, nil
	}

	// 2. Conectar a MongoDB
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return events.APIGatewayProxyResponse{StatusCode: 500, Headers: corsHeaders(), Body: `{"error": "Falta MONGODB_URI"}`}, nil
	}
	
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Headers: corsHeaders(), Body: fmt.Sprintf(`{"error": "%v"}`, err)}, nil
	}
	defer client.Disconnect(ctx)

	db := client.Database("personalbarber")
	productsColl := db.Collection("products")

	// 3. Extraer IDs de productos y buscar precios oficiales
	var ids []string
	for _, reqItem := range body.Items {
		ids = append(ids, reqItem.ID)
	}

	filter := bson.M{"id": bson.M{"$in": ids}}
	cursor, err := productsColl.Find(ctx, filter)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Headers: corsHeaders(), Body: `{"error": "Fallo al consultar BBDD"}`}, nil
	}
	defer cursor.Close(ctx)

	var foundProducts []Product
	if err := cursor.All(ctx, &foundProducts); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Headers: corsHeaders(), Body: `{"error": "Error leyendo productos"}`}, nil
	}

	// 4. Crear el cálculo cruzado
	var total float64 = 0
	var finalItems []bson.M

	for _, payloadItem := range body.Items {
		var matched *Product
		for _, fp := range foundProducts {
			if fp.ID == payloadItem.ID {
				matched = &fp
				break
			}
		}

		if matched != nil && payloadItem.Qty > 0 {
			price := parsePrice(matched.Price)
			subtotal := price * float64(payloadItem.Qty)
			total += subtotal

			finalItems = append(finalItems, bson.M{
				"id":    matched.ID,
				"name":  matched.Name,
				"qty":   payloadItem.Qty,
				"price": price,
				"subtotal": subtotal,
			})
		}
	}

	// Si el total es 0, hubo intento de hackear con IDs falsos
	if total == 0 {
		return events.APIGatewayProxyResponse{StatusCode: 400, Headers: corsHeaders(), Body: `{"error": "Productos inválidos o agotados"}`}, nil
	}

	// 4.1 Calcular costo de envío SIEMPRE en el servidor — NUNCA confiar en el frontend
	shippingCost := calculateShippingCost(body.ShippingMethod)
	subtotal := total
	grandTotal := subtotal + shippingCost

	// 5. Crear Orden con ID criptográficamente seguro
	orderID := generateOrderID()

	newOrder := Order{
		ID:             orderID,
		Customer:       body.Customer,
		Items:          finalItems,
		Subtotal:       subtotal,
		SubtotalFormat: formatNumberIntl(subtotal),
		ShippingCost:   shippingCost,
		ShippingFormat: formatNumberIntl(shippingCost),
		ShippingMethod: body.ShippingMethod,
		Total:          grandTotal,
		TotalFormat:    formatNumberIntl(grandTotal),
		PaymentMethod:  body.PaymentMethod,
		Status:         "PENDING",
		ClientIP:       getClientIP(request),
		CreatedAt:      time.Now(),
	}

	// 6. Guardar en Colección
	ordersColl := db.Collection("orders")
	_, err = ordersColl.InsertOne(ctx, newOrder)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Headers: corsHeaders(), Body: `{"error": "No se guardó la orden"}`}, nil
	}

	// 7. Retornar éxito con total recalculado
	resBytes, _ := json.Marshal(map[string]interface{}{
		"ok":    true,
		"order": newOrder,
	})

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    corsHeaders(),
		Body:       string(resBytes),
	}, nil
}

func main() {
	lambda.Start(handler)
}
