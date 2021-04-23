package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (h AppRouter) getComponent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	c := vars["component"]

	example, err := h.Doc.GetComponentExample(c)
	if err != nil {
		if err.Error() == "NotFound" {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
	}
	json.NewEncoder(w).Encode(example)
}
