package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/resend/resend-go/v2"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Reservation struct {
	Nombre      string    `json:"nombre" bson:"nombre"`
	Telefono    string    `json:"telefono" bson:"telefono"`
	Servicio    string    `json:"servicio" bson:"servicio"`
	FechaRaw    string    `json:"fechaRaw" bson:"fechaRaw"`
	HoraRaw     string    `json:"horaRaw" bson:"horaRaw"`
	Direccion   string    `json:"direccion" bson:"direccion"`
	WhatsappUrl string    `json:"whatsappUrl" bson:"whatsappUrl"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Solo aceptamos POST
	if request.HTTPMethod != "POST" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusMethodNotAllowed,
			Body:       `{"error": "Method not allowed"}`,
		}, nil
	}

	// Parsear el body
	var res Reservation
	err := json.Unmarshal([]byte(request.Body), &res)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       `{"error": "Invalid request body"}`,
		}, nil
	}

	// Sanitización y Validación de Longitud Estricta (Anti Spam/Dos Data)
	if len(res.Nombre) > 100 || len(res.Telefono) > 30 || len(res.Servicio) > 100 || 
	   len(res.Direccion) > 200 || len(res.FechaRaw) > 50 || len(res.HoraRaw) > 50 {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest, // 400 Bad Request
			Body:       `{"error": "Payload excedió longitud máxima permitida por seguridad"}`,
		}, nil
	}

	res.CreatedAt = time.Now()

	// Conectar a MongoDB
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       `{"error": "Database configuration missing"}`,
		}, nil
	}

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf(`{"error": "Failed to connect to DB: %v"}`, err),
		}, nil
	}
	defer client.Disconnect(ctx)

	// Ping a la base de datos para verificar conexión
	err = client.Ping(ctx, nil)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf(`{"error": "Failed to ping DB: %v"}`, err),
		}, nil
	}

	collection := client.Database("personalbarber").Collection("reservations")

	// Verificar si ya existe una reserva para la misma fecha y hora
	var existing Reservation
	filter := map[string]string{
		"fechaRaw": res.FechaRaw,
		"horaRaw":  res.HoraRaw,
	}
	err = collection.FindOne(ctx, filter).Decode(&existing)
	if err == nil {
		// Si no hay error, significa que encontró una reserva coincidente
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusConflict, // 409 Conflict
			Body:       `{"error": "Este turno ya ha sido reservado por otra persona. Por favor elige otro."}`,
		}, nil
	} else if err != mongo.ErrNoDocuments {
		// Si hubo un error que no sea "No se encontró el documento"
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf(`{"error": "Error al verificar disponibilidad: %v"}`, err),
		}, nil
	}

	// Si no existe, procedemos a insertar
	_, err = collection.InsertOne(ctx, res)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf(`{"error": "Failed to save reservation: %v"}`, err),
		}, nil
	}

	// === ENVÍO DE NOTIFICACIÓN POR CORREO (RESEND) ===
	resendKey := os.Getenv("RESEND_API_KEY")
	if resendKey != "" {
		clienteResend := resend.NewClient(resendKey)
		
		htmlContent := fmt.Sprintf(`
			<h2>¡Nueva Reserva Recibida! 💈</h2>
			<ul>
				<li><strong>Cliente:</strong> %s</li>
				<li><strong>Servicio:</strong> %s</li>
				<li><strong>Fecha:</strong> %s</li>
				<li><strong>Hora:</strong> %s</li>
				<li><strong>Teléfono:</strong> %s</li>
				<li><strong>Dirección:</strong> %s</li>
			</ul>
			<p><a href="%s">Abrir en WhatsApp</a></p>
		`, res.Nombre, res.Servicio, res.FechaRaw, res.HoraRaw, res.Telefono, res.Direccion, res.WhatsappUrl)

		parametrosCorreo := &resend.SendEmailRequest{
			From:    "onboarding@resend.dev",
			To:      []string{"jhonechavarria0506@gmail.com"},
			Subject: "💈 ¡Nueva Reserva de " + res.Nombre + "!",
			Html:    htmlContent,
		}

		_, errResend := clienteResend.Emails.Send(parametrosCorreo)
		if errResend != nil {
			// Solo logueamos el error para no afectar la respuesta exitosa de la reserva al usuario
			fmt.Printf("Error al enviar correo con Resend: %v\n", errResend)
		} else {
			fmt.Println("Correo de notificación enviado con éxito")
		}
	} else {
		fmt.Println("No se encontró RESEND_API_KEY. Saltando el envío de correo.")
	}
	// =================================================

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: `{"message": "Reserva guardada con éxito en MongoDB Atlas!"}`,
	}, nil
}

func main() {
	lambda.Start(handler)
}
