package handler

import (
	"net/http"

	"github.com/febrielven/go-crudapi/db"
	"github.com/go-chi/chi"
)

var dbInstance db.Database

func NewHandler(db db.Database) http.Handler {
	router := chi.NewRouter()
	dbInstance = db
	router.Me
}
