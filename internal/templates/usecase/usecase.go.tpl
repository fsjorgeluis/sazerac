package usecases

import "{{ .Module }}/internal/domain/entities"
import "{{ .Module }}/internal/domain/interfaces"

type {{ .Name }}UseCase struct {
    Repo interfaces.{{ .Entity }}Repository
}

func New{{ .Name }}UseCase(repo interfaces.{{ .Entity }}Repository) *{{ .Name }}UseCase {
    return &{{ .Name }}UseCase{Repo: repo}
}

func (uc *{{ .Name }}UseCase) Execute(input {{ .Name }}Input) (*entities.{{ .Entity }}, error) {
    // TODO: business logic here
    return nil, nil
}

type {{ .Name }}Input struct {
    // TODO: add definition
}