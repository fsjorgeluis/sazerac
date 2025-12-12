package di

import (
	{{- if eq .Features.Database "mysql" }}
	"database/sql"

	"{{ .Module }}/infrastructure/database/mysql"
	{{- else if eq .Features.Database "none" }}
	"{{ .Module }}/infrastructure/database/inmemory"
	{{- end }}
	"{{ .Module }}/internal/handlers"
	"{{ .Module }}/internal/usecases"
)

// Container holds all dependencies
type Container struct {
	{{- if eq .Features.Database "mysql" }}
	DB     *sql.DB
	{{- end }}
	{{ .UseCase }}Handler *handlers.{{ .UseCase }}Handler
}

// NewContainer initializes all dependencies and returns a Container
func NewContainer() (*Container, error) {
	{{- if eq .Features.Database "mysql" }}
	// Initialize database connection (optional for demo)
	// In production, you would initialize a real database connection here
	var db *sql.DB = nil

	// Initialize repository with nil DB (for demo purposes)
	// In production, you would pass a real database connection
	{{ .Entity }}Repo := mysql.New{{ .Entity }}MySQLRepo(db)
	{{- else }}
	// Using in-memory repository
	{{ .Entity }}Repo := inmemory.New{{ .Entity }}InMemoryRepo()
	{{- end }}

	// Initialize use case
	{{ .UseCase }}UC := usecases.New{{ .UseCase }}UseCase({{ .Entity }}Repo)

	// Initialize handler
	{{ .UseCase }}Handler := handlers.New{{ .UseCase }}Handler({{ .UseCase }}UC)

	return &Container{
		{{- if eq .Features.Database "mysql" }}
		DB:     db,
		{{- end }}
		{{ .UseCase }}Handler: {{ .UseCase }}Handler,
	}, nil
}

{{- if eq .Features.Database "mysql" }}
// Close closes all connections
func (c *Container) Close() error {
	if c.DB != nil {
		return c.DB.Close()
	}
	return nil
}
{{- end }}

