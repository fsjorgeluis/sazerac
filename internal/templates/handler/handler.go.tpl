package handlers

import (
	"fmt"

	"{{ .Module }}/internal/usecases"
)

type {{ .Name }}Handler struct {
	UC *usecases.{{ .UseCase }}UseCase
}

func New{{ .Name }}Handler(uc *usecases.{{ .UseCase }}UseCase) *{{ .Name }}Handler {
	return &{{ .Name }}Handler{UC: uc}
}

// Run executes the use case and displays the result
func (h *{{ .Name }}Handler) Run() error {
	input := usecases.{{ .UseCase }}Input{}
	
	entity, err := h.UC.Execute(input)
	if err != nil {
		return fmt.Errorf("failed to execute use case: %w", err)
	}
	
	fmt.Printf("Have a good drink! 🥃\n")
	fmt.Printf("Entity created: ID=%s, Name=%s\n", entity.ID, entity.Name)
	return nil
}