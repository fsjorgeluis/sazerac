package repository

import (
	"context"

	"{{ .Module }}/internal/domain/entities"
)

type {{ .Entity }}Repository interface {
    Save(ctx context.Context, e *entities.{{ .Entity }}) error
    FindByID(ctx context.Context, id string) (*entities.{{ .Entity }}, error)
}