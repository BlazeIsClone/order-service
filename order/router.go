package order

import (
	"net/http"
)

func HandleRoutes(router *http.ServeMux) {
	handler := &Handler{}

	router.HandleFunc("POST /orders", handler.Create)
	router.HandleFunc("PUT /orders/{id}", handler.UpdateByID)
	router.HandleFunc("GET /orders/{id}", handler.FindByID)
	router.HandleFunc("DELETE /orders/{id}", handler.DeleteByID)
}
