package router

import (
	"encoding/json"
	"net/http"
)

func (h AppRouter) getOperations(w http.ResponseWriter, r *http.Request) {
	tag := r.URL.Query().Get("tag")
	ops := h.Doc.GetOperationsByTag(tag)
	if len(ops) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(ops)
}
