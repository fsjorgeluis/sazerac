package handlers

import (
	"encoding/json"
	"net/http"
	"{{ .Module }}/internal/domain/usecases"
)

type {{ .Name }}Handler struct {
UC *usecases.{{ .UseCase }}UseCase
}

func New{{ .Name }}Handler(uc *usecases.{{ .UseCase }}UseCase) *{{ .Name }}Handler {
return &{{ .Name }}Handler{UC: uc}
}

func (h *{{ .Name }}Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
var input usecases.{{ .UseCase }}Input
if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
http.Error(w, err.Error(), http.StatusBadRequest)
return
}

result, err := h.UC.Execute(input)
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}

json.NewEncoder(w).Encode(result)
}