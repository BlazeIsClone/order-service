package product

import (
	"net/http"
)

func Routes(router *http.ServeMux) {
	handler := &Handler{}

	router.HandleFunc("POST /products", handler.Create)
	router.HandleFunc("GET /products/{id}", handler.FindByID)
	router.HandleFunc("PUT /products/{id}", handler.UpdateByID)
	router.HandleFunc("DELETE /products/{id}", handler.DeleteByID)
}
