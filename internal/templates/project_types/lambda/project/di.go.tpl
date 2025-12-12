package di

import (
	"context"
	{{- if eq .Features.Database "dynamodb" }}
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	{{- else if eq .Features.Database "mysql-rds" }}
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	{{- end }}

	{{- if eq .Features.Database "dynamodb" }}
	dynamoRepo "{{ .Module }}/infrastructure/database/dynamodb"
	{{- else if eq .Features.Database "mysql-rds" }}
	"{{ .Module }}/infrastructure/database/mysql"
	{{- else if eq .Features.Database "none" }}
	inmemoryRepo "{{ .Module }}/infrastructure/database/inmemory"
	{{- end }}
	"{{ .Module }}/internal/handlers"
	"{{ .Module }}/internal/usecases"
)

// Container holds all dependencies for the Lambda function
type Container struct {
	{{- if ne .Features.Database "none" }}
	{{- if eq .Features.Database "dynamodb" }}
	DynamoClient *dynamodb.Client
	{{- else if eq .Features.Database "mysql-rds" }}
	DB *sql.DB
	{{- end }}
	{{- end }}
	{{ .UseCase }}Handler *handlers.{{ .UseCase }}Handler
}

// NewContainer initializes all dependencies
func NewContainer(ctx context.Context) (*Container, error) {
	{{- if eq .Features.Database "dynamodb" }}
	// Initialize AWS config
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	dynamoClient := dynamodb.NewFromConfig(cfg)
	tableName := os.Getenv("TABLE_NAME")
	if tableName == "" {
		tableName = "{{ .ProjectName }}-table"
	}

	// Initialize repository
	{{ .Entity }}Repo := dynamoRepo.New{{ .Entity }}DynamoDBRepo(dynamoClient, tableName)
	{{- else if eq .Features.Database "mysql-rds" }}
	// Initialize database connection
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "user:password@tcp(localhost:3306)/dbname"
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Initialize repository
	{{ .Entity }}Repo := mysql.New{{ .Entity }}MySQLRepo(db)
	{{- else if eq .Features.Database "none" }}
	// Using in-memory repository
	{{ .Entity }}Repo := inmemoryRepo.New{{ .Entity }}InMemoryRepo()
	{{- end }}

	// Initialize use case
	{{ .UseCase }}UC := usecases.New{{ .UseCase }}UseCase({{ .Entity }}Repo)

	// Initialize handler
	{{ .UseCase }}Handler := handlers.New{{ .UseCase }}Handler({{ .UseCase }}UC)

	return &Container{
		{{- if eq .Features.Database "dynamodb" }}
		DynamoClient: dynamoClient,
		{{- else if eq .Features.Database "mysql-rds" }}
		DB: db,
		{{- end }}
		{{ .UseCase }}Handler: {{ .UseCase }}Handler,
	}, nil
}

// Close cleans up resources
func (c *Container) Close() error {
	{{- if eq .Features.Database "mysql-rds" }}
	if c.DB != nil {
		return c.DB.Close()
	}
	{{- end }}
	return nil
}
