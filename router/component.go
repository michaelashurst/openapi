package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/michaelashurst/openapi/document/component"
)

func (h AppRouter) getComponent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	c := vars["component"]

	for key, s := range h.Doc.Components.Schemas {
		if key == c {
			example := component.ConvertSchemaToExample(s)
			json.NewEncoder(w).Encode(example)
			break
		}
	}
}
