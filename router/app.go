package router

import (
	"os"

	"github.com/gorilla/mux"
	"github.com/michaelashurst/openapi/document"
)

type AppRouter struct {
	Doc *document.Document
}

func NewAppRouter(path string) AppRouter {
	doc := document.NewDocument(os.Args[1])
	return AppRouter{Doc: doc}
}

func (h AppRouter) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/api/paths", h.getPaths).Methods("GET")
	router.HandleFunc("/api/operations", h.getOperations).Methods("GET")

	router.HandleFunc("/api/servers", h.getServers).Methods("GET")

	router.HandleFunc("/api/component/{component}", h.getComponent).Methods("GET")

	// Tags
	router.HandleFunc("/api/tags", h.getTags).Methods("GET")
	router.HandleFunc("/api/tag/{tag}", h.getTag).Methods("GET")

	// Info
	router.HandleFunc("/api/info", h.getInfo).Methods("GET")
	router.HandleFunc("/api/info", h.postInfo).Methods("POST")
}
