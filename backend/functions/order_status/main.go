package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// ──────────────────────────────────────────
// Consulta el estado actual de una orden.
// Usado por el frontend para hacer polling después del pago.
//
// GET /api/order_status?id=ORD-12345
// ──────────────────────────────────────────

type OrderResponse struct {
	ID             string  `bson:"id" json:"id"`
	Status         string  `bson:"status" json:"status"`
	Total          float64 `bson:"total" json:"total"`
	TotalFormat    string  `bson:"total_format" json:"total_format"`
	ShippingMethod string  `bson:"shippingMethod" json:"shippingMethod"`
	PaymentMethod  string  `bson:"paymentMethod" json:"paymentMethod"`
	// Campos Wompi
	WompiTransactionID string `bson:"wompiTransactionId,omitempty" json:"wompiTransactionId,omitempty"`
	WompiStatus        string `bson:"wompiStatus,omitempty" json:"wompiStatus,omitempty"`
}

func corsHeaders() map[string]string {
	return map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Headers": "Content-Type",
		"Access-Control-Allow-Methods": "GET, OPTIONS",
		"Content-Type":                 "application/json",
	}
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Preflight CORS
	if request.HTTPMethod == "OPTIONS" {
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Headers:    corsHeaders(),
			Body:       "OK",
		}, nil
	}

	// ── 1. Obtener el ID de la orden ──
	orderID := request.QueryStringParameters["id"]
	if orderID == "" {
		body, _ := json.Marshal(map[string]interface{}{
			"ok":    false,
			"error": "El parámetro 'id' es requerido",
		})
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers:    corsHeaders(),
			Body:       string(body),
		}, nil
	}

	// Validar formato del ID (prevenir inyecciones)
	if len(orderID) > 20 {
		body, _ := json.Marshal(map[string]interface{}{
			"ok":    false,
			"error": "ID inválido",
		})
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers:    corsHeaders(),
			Body:       string(body),
		}, nil
	}

	// ── 2. Conectar a MongoDB ──
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers:    corsHeaders(),
			Body:       `{"ok": false, "error": "config"}`,
		}, nil
	}

	client, err := mongo.Connect(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers:    corsHeaders(),
			Body:       `{"ok": false, "error": "db_connect"}`,
		}, nil
	}
	defer client.Disconnect(ctx)

	// ── 3. Buscar la orden ──
	db := client.Database("personalbarber")
	var order OrderResponse
	err = db.Collection("orders").FindOne(ctx, bson.M{"id": orderID}).Decode(&order)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			body, _ := json.Marshal(map[string]interface{}{
				"ok":    false,
				"error": "Orden no encontrada",
			})
			return events.APIGatewayProxyResponse{
				StatusCode: 404,
				Headers:    corsHeaders(),
				Body:       string(body),
			}, nil
		}
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers:    corsHeaders(),
			Body:       fmt.Sprintf(`{"ok": false, "error": "%v"}`, err),
		}, nil
	}

	// ── 4. Responder ──
	body, _ := json.Marshal(map[string]interface{}{
		"ok":    true,
		"order": order,
	})

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    corsHeaders(),
		Body:       string(body),
	}, nil
}

func main() {
	lambda.Start(handler)
}
