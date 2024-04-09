package product

import (
	"net/http"
	"os"
)

func Routes(router *http.ServeMux) {
	handler := &Handler{}

	basePath := os.Getenv("BASE_PATH")

	router.HandleFunc("POST "+basePath+"/products", handler.Create)
	router.HandleFunc("GET "+basePath+"/products/{id}", handler.FindByID)
	router.HandleFunc("PUT "+basePath+"/products/{id}", handler.UpdateByID)
	router.HandleFunc("DELETE "+basePath+"/products/{id}", handler.DeleteByID)
}
