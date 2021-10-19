package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/chi"

	"github.com/ianyong/todo-backend/internal/adapters/infrastructure/database"
	"github.com/ianyong/todo-backend/internal/adapters/userinterface/router"
	"github.com/ianyong/todo-backend/internal/auth"
	"github.com/ianyong/todo-backend/internal/config"
	"github.com/ianyong/todo-backend/internal/services"
)

var l *chiadapter.ChiLambda

// main is the entry point for the AWS Lambda server.
func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v\n", err)
	}

	db, err := database.SetUp(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v\n", err)
	}

	jwtManager, err := auth.NewJWTManager(cfg.SecretKey, nil)
	if err != nil {
		log.Fatalf("failed to create JWT manager: %v\n", err)
	}

	s := services.SetUp(db, jwtManager)
	r := router.SetUp(s, cfg)
	l = chiadapter.New(r)

	lambda.Start(lambdaHandler)
}

// lambdaHandler acts as a wrapper around the API, translating requests and responses into the format
// expected by AWS Lambda.
func lambdaHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return l.ProxyWithContext(ctx, req)
}
