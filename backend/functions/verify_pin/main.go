package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/JohnmaDev/PersonalBarber/backend/pkg/auth"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	headers := map[string]string{
		"Content-Type":                "application/json",
		"Access-Control-Allow-Origin": "*",
	}

	if request.HTTPMethod == "OPTIONS" {
		return events.APIGatewayProxyResponse{StatusCode: 200, Headers: headers}, nil
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers:    headers,
			Body:       `{"ok": false, "error": "Missing MONGODB_URI"}`,
		}, nil
	}

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers:    headers,
			Body:       fmt.Sprintf(`{"ok": false, "error": "%v"}`, err),
		}, nil
	}
	defer client.Disconnect(ctx)

	ok, err := auth.VerifyTokenWithRateLimit(ctx, request, client)
	if !ok {
		msg := "PIN incorrecto o IP Bloqueada"
		if err != nil {
			msg = err.Error()
		}
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusUnauthorized,
			Headers:    headers,
			Body:       fmt.Sprintf(`{"ok": false, "error": "%s"}`, msg),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    headers,
		Body:       `{"ok": true, "message": "Autenticado"}`,
	}, nil
}

func main() {
	lambda.Start(handler)
}
