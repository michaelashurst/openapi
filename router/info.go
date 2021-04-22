package router

import (
	"encoding/json"
	"net/http"

	"github.com/michaelashurst/openapi/document/info"
)

func (h AppRouter) getInfo(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.Doc.Info)
}

func (h AppRouter) postInfo(w http.ResponseWriter, r *http.Request) {
	var info info.Info
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.Doc.UpdateInfo(info)
}
