package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (h AppRouter) getComponent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	c := vars["component"]

	for key, s := range h.Doc.Components.Schemas {
		if key == c {
			example := s.GenerateExample()
			json.NewEncoder(w).Encode(example)
			break
		}
	}
}
