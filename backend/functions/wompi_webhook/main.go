package main

import (
	"context"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// ──────────────────────────────────────────
// Webhook de Wompi — recibe eventos transaction.updated
//
// Verifica la firma del evento, actualiza la orden en MongoDB,
// y descuenta stock cuando el pago es aprobado.
//
// POST /api/wompi_webhook
// ──────────────────────────────────────────

// ── Tipos del payload de Wompi ──

type WebhookPayload struct {
	Event     string        `json:"event"`
	Data      WebhookData   `json:"data"`
	Signature WebhookSig    `json:"signature"`
	Timestamp int64         `json:"timestamp"`
	SentAt    string        `json:"sent_at"`
}

type WebhookData struct {
	Transaction WompiTransaction `json:"transaction"`
}

type WompiTransaction struct {
	ID            string `json:"id"`
	Status        string `json:"status"`
	Reference     string `json:"reference"`
	AmountInCents int64  `json:"amount_in_cents"`
	Currency      string `json:"currency"`
	PaymentMethod interface{} `json:"payment_method"`
	CreatedAt     string `json:"created_at"`
	FinalizedAt   string `json:"finalized_at"`
}

type WebhookSig struct {
	Checksum   string   `json:"checksum"`
	Properties []string `json:"properties"`
}

// ── Tipo para items de la orden en MongoDB ──

type OrderItem struct {
	ID  interface{} `bson:"id"`
	Qty int         `bson:"qty"`
}

type StoredOrder struct {
	ID     string      `bson:"id"`
	Status string      `bson:"status"`
	Items  []OrderItem `bson:"items"`
}

// corsHeaders devuelve las cabeceras CORS
func corsHeaders() map[string]string {
	return map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Headers": "Content-Type, X-Event-Checksum",
		"Access-Control-Allow-Methods": "POST, OPTIONS",
		"Content-Type":                 "application/json",
	}
}

// getNestedValue navega las propiedades del payload siguiendo la notación de punto
// Ejemplo: "transaction.id" → data.transaction.id
func getNestedValue(data WebhookData, path string) string {
	parts := strings.Split(path, ".")
	if len(parts) < 2 || parts[0] != "transaction" {
		return ""
	}

	tx := data.Transaction
	switch parts[1] {
	case "id":
		return tx.ID
	case "status":
		return tx.Status
	case "amount_in_cents":
		return fmt.Sprintf("%d", tx.AmountInCents)
	case "reference":
		return tx.Reference
	case "currency":
		return tx.Currency
	default:
		return ""
	}
}

// verifySignature valida la firma del webhook de Wompi
func verifySignature(payload WebhookPayload, eventsSecret string) bool {
	// Concatenar los valores de las propiedades en orden
	var concat string
	for _, prop := range payload.Signature.Properties {
		concat += getNestedValue(payload.Data, prop)
	}
	// Agregar timestamp y secreto
	concat += fmt.Sprintf("%d", payload.Timestamp)
	concat += eventsSecret

	// SHA256
	hash := sha256.Sum256([]byte(concat))
	calculated := hex.EncodeToString(hash[:])

	// Comparación timing-safe para prevenir timing attacks
	return subtle.ConstantTimeCompare([]byte(calculated), []byte(payload.Signature.Checksum)) == 1
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

	// ── 1. Validar configuración ──
	mongoURI := os.Getenv("MONGODB_URI")
	eventsSecret := os.Getenv("WOMPI_EVENTS_SECRET")

	if mongoURI == "" || eventsSecret == "" {
		fmt.Println("[WEBHOOK ERROR] Variables de entorno faltantes")
		// Responder 200 para que Wompi no reintente infinitamente
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Headers:    corsHeaders(),
			Body:       `{"ok": false, "error": "config"}`,
		}, nil
	}

	// ── 2. Parsear el payload ──
	var payload WebhookPayload
	if err := json.Unmarshal([]byte(request.Body), &payload); err != nil {
		fmt.Printf("[WEBHOOK ERROR] JSON inválido: %v\n", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Headers:    corsHeaders(),
			Body:       `{"ok": false, "error": "invalid_json"}`,
		}, nil
	}

	// ── 3. Verificar la firma ──
	if !verifySignature(payload, eventsSecret) {
		fmt.Printf("[WEBHOOK SECURITY] Firma inválida para evento %s ref=%s\n",
			payload.Event, payload.Data.Transaction.Reference)
		return events.APIGatewayProxyResponse{
			StatusCode: 401,
			Headers:    corsHeaders(),
			Body:       `{"ok": false, "error": "invalid_signature"}`,
		}, nil
	}

	// ── 4. Solo procesar transaction.updated ──
	if payload.Event != "transaction.updated" {
		fmt.Printf("[WEBHOOK] Evento ignorado: %s\n", payload.Event)
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Headers:    corsHeaders(),
			Body:       `{"ok": true, "action": "ignored"}`,
		}, nil
	}

	tx := payload.Data.Transaction
	fmt.Printf("[WEBHOOK] Procesando transacción %s ref=%s status=%s amount=%d\n",
		tx.ID, tx.Reference, tx.Status, tx.AmountInCents)

	// ── 5. Conectar a MongoDB ──
	client, err := mongo.Connect(options.Client().ApplyURI(mongoURI))
	if err != nil {
		fmt.Printf("[WEBHOOK ERROR] MongoDB connect: %v\n", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Headers:    corsHeaders(),
			Body:       `{"ok": false, "error": "db_connect"}`,
		}, nil
	}
	defer client.Disconnect(ctx)

	db := client.Database("personalbarber")
	ordersColl := db.Collection("orders")

	// ── 6. Buscar la orden por referencia ──
	var order StoredOrder
	err = ordersColl.FindOne(ctx, bson.M{
		"id": tx.Reference,
	}).Decode(&order)

	if err != nil {
		fmt.Printf("[WEBHOOK WARN] Orden no encontrada: %s\n", tx.Reference)
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Headers:    corsHeaders(),
			Body:       `{"ok": true, "action": "order_not_found"}`,
		}, nil
	}

	// Evitar reprocesar órdenes ya finalizadas
	if order.Status == "APPROVED" || order.Status == "DECLINED" || order.Status == "VOIDED" {
		fmt.Printf("[WEBHOOK] Orden %s ya en estado final: %s (ignorando)\n", order.ID, order.Status)
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Headers:    corsHeaders(),
			Body:       `{"ok": true, "action": "already_processed"}`,
		}, nil
	}

	// ── 7. Actualizar la orden según el estado de la transacción ──
	now := time.Now()

	updateFields := bson.M{
		"wompiTransactionId": tx.ID,
		"wompiStatus":        tx.Status,
		"updatedAt":          now,
	}

	switch tx.Status {
	case "APPROVED":
		updateFields["status"] = "APPROVED"
		updateFields["paidAt"] = now
		updateFields["paymentConfirmed"] = true
	case "DECLINED":
		updateFields["status"] = "DECLINED"
		updateFields["declinedAt"] = now
	case "VOIDED":
		updateFields["status"] = "VOIDED"
		updateFields["voidedAt"] = now
	case "ERROR":
		updateFields["status"] = "ERROR"
		updateFields["errorAt"] = now
	default:
		// Estados intermedios (PENDING, etc.) — solo guardar el ID de transacción
		updateFields["status"] = tx.Status
	}

	_, err = ordersColl.UpdateOne(ctx,
		bson.M{"id": order.ID},
		bson.M{"$set": updateFields},
	)
	if err != nil {
		fmt.Printf("[WEBHOOK ERROR] No se pudo actualizar orden %s: %v\n", order.ID, err)
	}

	// ── 8. Si fue aprobada, descontar stock ──
	if tx.Status == "APPROVED" {
		productsColl := db.Collection("products")
		for _, item := range order.Items {
			if item.Qty > 0 {
				_, err := productsColl.UpdateOne(ctx,
					bson.M{"id": item.ID},
					bson.M{"$inc": bson.M{"stock": -item.Qty}},
				)
				if err != nil {
					fmt.Printf("[WEBHOOK WARN] Error descontando stock de producto %v: %v\n", item.ID, err)
				}
			}
		}
		fmt.Printf("[WEBHOOK OK] Orden %s APROBADA — stock descontado\n", order.ID)
	} else {
		fmt.Printf("[WEBHOOK OK] Orden %s actualizada a %s\n", order.ID, tx.Status)
	}

	// ── 9. Siempre responder 200 a Wompi ──
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    corsHeaders(),
		Body:       fmt.Sprintf(`{"ok": true, "action": "processed", "status": "%s"}`, tx.Status),
	}, nil
}

func main() {
	lambda.Start(handler)
}
