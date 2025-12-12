package mysql

import (
	"context"
	"database/sql"

	"{{ .Module }}/internal/domain/entities"
	{{- if .Features.ErrorHandling }}
	"{{ .Module }}/internal/domain/errors"
	{{- end }}
	"{{ .Module }}/internal/repository"
)

type {{ .Entity }}MySQLRepo struct {
	DB *sql.DB
}

func New{{ .Entity }}MySQLRepo(db *sql.DB) repository.{{ .Entity }}Repository {
	return &{{ .Entity }}MySQLRepo{DB: db}
}

func (r *{{ .Entity }}MySQLRepo) Save(ctx context.Context, e *entities.{{ .Entity }}) error {
	query := `INSERT INTO {{ .Entity | ToLower }}s (id, name) VALUES (?, ?)
		ON DUPLICATE KEY UPDATE name = VALUES(name)`
	
	_, err := r.DB.ExecContext(ctx, query, e.ID, e.Name)
	{{- if .Features.ErrorHandling }}
	if err != nil {
		return errors.ErrInternalServer
	}
	{{- end }}
	
	return err
}

func (r *{{ .Entity }}MySQLRepo) FindByID(ctx context.Context, id string) (*entities.{{ .Entity }}, error) {
	query := `SELECT id, name FROM {{ .Entity | ToLower }}s WHERE id = ?`
	
	var entity entities.{{ .Entity }}
	err := r.DB.QueryRowContext(ctx, query, id).Scan(&entity.ID, &entity.Name)
	
	if err == sql.ErrNoRows {
		{{- if .Features.ErrorHandling }}
		return nil, errors.NewNotFoundError("{{ .Entity }}")
		{{- else }}
		return nil, nil
		{{- end }}
	}
	
	if err != nil {
		{{- if .Features.ErrorHandling }}
		return nil, errors.ErrInternalServer
		{{- else }}
		return nil, err
		{{- end }}
	}
	
	return &entity, nil
}
