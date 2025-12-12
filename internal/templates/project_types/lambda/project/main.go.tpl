package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

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
	{{- if .UseCase }}
	return container.{{ .UseCase }}Handler.Handle(ctx, request)
	{{- else }}
	// TODO: Update this handler after running 'sazerac make all <Entity> <UseCase>'
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       `{"message": "Lambda function initialized. Run 'sazerac make all <Entity> <UseCase>' to generate handlers."}`,
	}, nil
	{{- end }}
}

func main() {
	lambda.Start(handler)
}
