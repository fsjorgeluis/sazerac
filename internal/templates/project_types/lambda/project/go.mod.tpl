module {{ .Module }}

go 1.21

{{- if eq .Features.Database "dynamodb" }}
require (
	github.com/aws/aws-lambda-go v1.41.0
	github.com/aws/aws-sdk-go-v2 v1.25.0
	github.com/aws/aws-sdk-go-v2/config v1.27.0
	github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue v1.13.0
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.30.0
)
{{- else if eq .Features.Database "mysql-rds" }}
require (
	github.com/aws/aws-lambda-go v1.41.0
	github.com/go-sql-driver/mysql v1.7.1
)
{{- else }}
require github.com/aws/aws-lambda-go v1.41.0
{{- end }}

// For local development, tell Go to use the current directory
replace {{ .Module }} => ./
