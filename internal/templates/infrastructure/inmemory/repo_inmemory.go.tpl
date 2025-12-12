package inmemory

import (
	"context"
	{{- if not .Features.ErrorHandling }}
	"fmt"
	{{- end }}
	"sync"

	"{{ .Module }}/internal/domain/entities"
	{{- if .Features.ErrorHandling }}
	"{{ .Module }}/internal/domain/errors"
	{{- end }}
	"{{ .Module }}/internal/repository"
)

type {{ .Entity }}InMemoryRepo struct {
	mu    sync.RWMutex
	store map[string]*entities.{{ .Entity }}
}

func New{{ .Entity }}InMemoryRepo() repository.{{ .Entity }}Repository {
	return &{{ .Entity }}InMemoryRepo{
		store: make(map[string]*entities.{{ .Entity }}),
	}
}

func (r *{{ .Entity }}InMemoryRepo) Save(ctx context.Context, e *entities.{{ .Entity }}) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	r.store[e.ID] = e
	return nil
}

func (r *{{ .Entity }}InMemoryRepo) FindByID(ctx context.Context, id string) (*entities.{{ .Entity }}, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	entity, ok := r.store[id]
	if !ok {
		{{- if .Features.ErrorHandling }}
		return nil, errors.NewNotFoundError("{{ .Entity }}")
		{{- else }}
		return nil, fmt.Errorf("{{ .Entity | ToLower }} not found")
		{{- end }}
	}
	
	return entity, nil
}
