package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	{{- if .Features.ErrorHandling }}
	"{{ .Module }}/internal/domain/errors"
	{{- end }}
	"{{ .Module }}/internal/usecases"
)

type {{ .Name }}Handler struct {
	UC *usecases.{{ .UseCase }}UseCase
}

func New{{ .Name }}Handler(uc *usecases.{{ .UseCase }}UseCase) *{{ .Name }}Handler {
	return &{{ .Name }}Handler{UC: uc}
}

// Handle processes Lambda requests from API Gateway
func (h *{{ .Name }}Handler) Handle(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var input usecases.{{ .UseCase }}Input
	if err := json.Unmarshal([]byte(request.Body), &input); err != nil {
		{{- if .Features.ErrorHandling }}
		return buildErrorResponse(errors.ErrBadRequest), nil
		{{- else }}
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       `{"error": "invalid request"}`,
			Headers:    map[string]string{"Content-Type": "application/json"},
		}, nil
		{{- end }}
	}

	result, err := h.UC.Execute(ctx, input)
	if err != nil {
		{{- if .Features.ErrorHandling }}
		return buildErrorResponse(err), nil
		{{- else }}
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       `{"error": "internal server error"}`,
			Headers:    map[string]string{"Content-Type": "application/json"},
		}, nil
		{{- end }}
	}

	body, _ := json.Marshal(result)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}

{{- if .Features.ErrorHandling }}

func buildErrorResponse(err error) events.APIGatewayProxyResponse {
	domainErr, ok := err.(*errors.DomainError)
	if !ok {
		domainErr = errors.ErrInternalServer
	}

	body, _ := json.Marshal(map[string]string{
		"code":    domainErr.Code,
		"message": domainErr.Message,
	})

	return events.APIGatewayProxyResponse{
		StatusCode: domainErr.HTTPStatus,
		Body:       string(body),
		Headers:    map[string]string{"Content-Type": "application/json"},
	}
}
{{- end }}
