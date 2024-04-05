package user

import (
	"net/http"
)

func LoadUserRoutes(router *http.ServeMux) {
	handler := &Handler{}

	router.HandleFunc("POST /user", handler.Create)
	router.HandleFunc("PUT /user/{id}", handler.UpdateByID)
	router.HandleFunc("GET /user/{id}", handler.FindByID)
	router.HandleFunc("DELETE /user/{id}", handler.DeleteByID)
}
