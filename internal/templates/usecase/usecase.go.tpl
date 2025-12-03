package usecases

import (
	"fmt"
	"math/rand"
	"time"

	"{{ .Module }}/internal/domain/entities"
	"{{ .Module }}/internal/repository"
)

type {{ .Name }}UseCase struct {
    Repo repository.{{ .Entity }}Repository
}

func New{{ .Name }}UseCase(repo repository.{{ .Entity }}Repository) *{{ .Name }}UseCase {
    return &{{ .Name }}UseCase{Repo: repo}
}

func (uc *{{ .Name }}UseCase) Execute(input {{ .Name }}Input) (*entities.{{ .Entity }}, error) {
    // TODO: business logic here
    
    // Generate random name for demo
    rand.Seed(time.Now().UnixNano())
    names := []string{"Alice", "Bob", "Charlie", "Diana", "Eve", "Frank", "Grace", "Henry"}
    randomName := names[rand.Intn(len(names))]
    
    entity := &entities.{{ .Entity }}{
        ID:   fmt.Sprintf("%d", time.Now().Unix()),
        Name: randomName,
    }
    
    // Save entity using repository
    if err := uc.Repo.Save(entity); err != nil {
        return nil, fmt.Errorf("failed to save entity: %w", err)
    }
    
    return entity, nil
}

type {{ .Name }}Input struct {
    // TODO: add definition
}