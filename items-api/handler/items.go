package handler

import (
	"context"
	"fmt"
)

var itemIDKey = "itemID"

func items(router chi.Router) {
	router.GET("/", getAllItems)
	router.POST("/", createItem)
	router.Route("/{itemId}", func(router chi.Router) {
		router.Use(ItemContext)
		router.GET("/", getItem)
		router.PUT("/", updateItem)
		router.Delete("/", deleteItem)
	})
}

func ItemContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		itemId := chi.URLParam(r, "itemId")
		if itemId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("Invalid item ID")))
		}
		ctx := context.WithValue(r.Context(), itemIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
