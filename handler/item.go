package handler

import (
	"net/http"

	"github.com/febrielven/go-crudapi/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// var itemIdKey = "itemID"

func items(router chi.Router) {
	router.Get("/", getAllItems)
	router.Post("/", createItem)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	item := &models.Item{}
	if err := render.Bind(r, item); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	if err := dbInstance.AddItem(item); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, item); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}

}

func getAllItems(w http.ResponseWriter, r *http.Request) {
	items, err := dbInstance.GetAllItems()
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}

	if err := render.Render(w, r, items); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
	}
}
