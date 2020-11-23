package handler

import (
	"net/http"

	"github.com/Neil-uli/Resful-go/items-api/db"
	"github.com/go-chi/chi"
)

var dbInstance db.Database

func NewHandler(db db.Database) http.Handler {
	router := chi.NewRouter()
	dbInstance = db
	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)
	router.Route("/items", items)
	return router
}
