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
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<style>
	body { margin: 0; padding: 0; background-color: #0f1014; font-family: 'Helvetica Neue', Helvetica, Arial, sans-serif; color: #ffffff; }
	.container { max-width: 600px; margin: 40px auto; background-color: #1a1c23; border-radius: 16px; overflow: hidden; box-shadow: 0 4px 20px rgba(0,0,0,0.5); border: 1px solid #2a2d36; }
	.header { background-color: #0f1014; padding: 30px 20px; text-align: center; border-bottom: 2px solid #39FF14; }
	.header h1 { margin: 0; color: #ffffff; font-size: 24px; font-weight: 800; letter-spacing: 1px; }
	.header span { color: #39FF14; }
	.content { padding: 40px 30px; }
	.greeting { font-size: 18px; margin-bottom: 25px; color: #e2e8f0; }
	.details-card { background-color: #242731; border-radius: 12px; padding: 25px; margin-bottom: 30px; border: 1px solid #2f3340; }
	.detail-item { margin-bottom: 15px; display: flex; flex-direction: column; }
	.detail-item:last-child { margin-bottom: 0; }
	.detail-label { font-size: 12px; text-transform: uppercase; letter-spacing: 1.5px; color: #94a3b8; margin-bottom: 4px; font-weight: 600; }
	.detail-value { font-size: 16px; color: #ffffff; font-weight: 500; }
	.btn-container { text-align: center; margin-top: 35px; }
	.btn { display: inline-block; background-color: #39FF14; color: #000000; text-decoration: none; padding: 14px 30px; border-radius: 50px; font-weight: 700; font-size: 16px; transition: all 0.3s ease; text-transform: uppercase; letter-spacing: 1px; box-shadow: 0 0 15px rgba(57, 255, 20, 0.4); }
	.btn:hover { background-color: #2ce60d; box-shadow: 0 0 25px rgba(57, 255, 20, 0.6); }
	.footer { text-align: center; padding: 20px; font-size: 12px; color: #64748b; background-color: #0f1014; border-top: 1px solid #1f222b; }
</style>
</head>
<body>
	<div class="container">
		<div class="header">
			<h1>PERSONAL <span>BARBER</span> 💈</h1>
		</div>
		<div class="content">
			<div class="greeting">¡Hola! Tienes una nueva reserva confirmada.</div>
			
			<div class="details-card">
				<div class="detail-item">
					<span class="detail-label">Cliente</span>
					<span class="detail-value">%s</span>
				</div>
				<div class="detail-item">
					<span class="detail-label">Servicio</span>
					<span class="detail-value">%s</span>
				</div>
				<div class="detail-item">
					<span class="detail-label">Fecha y Hora</span>
					<span class="detail-value">%s a las %s</span>
				</div>
				<div class="detail-item">
					<span class="detail-label">Teléfono</span>
					<span class="detail-value">%s</span>
				</div>
				<div class="detail-item">
					<span class="detail-label">Dirección</span>
					<span class="detail-value">%s</span>
				</div>
			</div>

			<div class="btn-container">
				<a href="%s" class="btn">Abrir chat en WhatsApp</a>
			</div>
		</div>
		<div class="footer">
			Notificación automática generada por el sistema de reservas.
		</div>
	</div>
</body>
</html>
		`, res.Nombre, res.Servicio, res.FechaRaw, res.HoraRaw, res.Telefono, res.Direccion, res.WhatsappUrl)

		parametrosCorreo := &resend.SendEmailRequest{
			From:    "Personal Barber <reservas@personalbarber.vip>",
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
