package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	{{- if eq .Features.Database "dynamodb" }}
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	{{- end }}

	"{{ .Module }}/cmd/lambda/di"
)

var container *di.Container

func init() {
	var err error
	container, err = di.NewContainer(context.Background())
	if err != nil {
		log.Fatalf("Failed to initialize container: %v", err)
	}
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return container.{{ .UseCase }}Handler.Handle(ctx, request)
}

func main() {
	lambda.Start(handler)
}
