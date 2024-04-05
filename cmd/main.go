package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/blazeisclone/order-service/domain/product"
)

func handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("order service reached\n"))
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", handle)
	product.Routes(router)

	os.Setenv("PORT", "3000")

	port := os.Getenv("PORT")

	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	fmt.Println("server listening on port:", port)

	server.ListenAndServe()
}
