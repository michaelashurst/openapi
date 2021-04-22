package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (h AppRouter) getTags(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.Doc.Tags)
}

func (h AppRouter) getTag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tag := vars["tag"]

	for _, t := range h.Doc.Tags {
		if t.Name == tag {
			json.NewEncoder(w).Encode(t)
			break
		}
	}
}
