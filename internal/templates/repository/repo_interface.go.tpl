package repository

import "{{ .Module }}/internal/domain/entities"

type {{ .Entity }}Repository interface {
    Save(e *entities.{{ .Entity }}) error
    FindByID(id string) (*entities.{{ .Entity }}, error)
}