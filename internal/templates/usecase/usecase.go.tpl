package usecases

import (
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
    return nil, nil
}

type {{ .Name }}Input struct {
    // TODO: add definition
}