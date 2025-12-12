package mappers

import "{{ .Module }}/internal/domain/entities"

func Map{{ .Entity }}FromDTO(dto any) (*entities.{{ .Entity }}, error) {
    // TODO: implement mapper here
    return nil, nil
}

func Map{{ .Entity }}ToDTO(e *entities.{{ .Entity }}) any {
    // TODO: implement mapper here
    return nil
}