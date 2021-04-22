package router

import (
	"encoding/json"
	"net/http"
)

func (h AppRouter) getPaths(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.Doc.Paths)
}
