package di

import (
	"database/sql"

	"{{ .Module }}/infrastructure/database/mysql"
	"{{ .Module }}/internal/handlers"
	"{{ .Module }}/internal/usecases"
)

// Container holds all dependencies
type Container struct {
	DB     *sql.DB
	{{ .UseCase }}Handler *handlers.{{ .UseCase }}Handler
}

// NewContainer initializes all dependencies and returns a Container
// For demo purposes, database connection is optional (nil is acceptable)
func NewContainer() (*Container, error) {
	// Initialize database connection (optional for demo)
	// In production, you would initialize a real database connection here
	var db *sql.DB = nil

	// Initialize repository with nil DB (for demo purposes)
	// In production, you would pass a real database connection
	{{ .Entity }}Repo := mysql.New{{ .Entity }}MySQLRepo(db)

	// Initialize use case
	{{ .UseCase }}UC := usecases.New{{ .UseCase }}UseCase({{ .Entity }}Repo)

	// Initialize handler
	{{ .UseCase }}Handler := handlers.New{{ .UseCase }}Handler({{ .UseCase }}UC)

	return &Container{
		DB:     db,
		{{ .UseCase }}Handler: {{ .UseCase }}Handler,
	}, nil
}

// Close closes all connections
func (c *Container) Close() error {
	if c.DB != nil {
		return c.DB.Close()
	}
	return nil
}

