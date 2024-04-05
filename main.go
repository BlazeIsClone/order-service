package main

import (
	"fmt"
	"net/http"

	"github.com/blazeisclone/order-service/order"
)

func handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("order service reached"))
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", handle)
	order.HandleRoutes(router)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server listening on port :8080")

	server.ListenAndServe()
}
