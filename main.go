package main

import (
	"fmt"
	"net/http"

	"github.com/blazeisclone/user-api-service/user"
)

func handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, from api..."))
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/api/v1/", handle)
	user.LoadUserRoutes(router)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server listening on port :8080")

	server.ListenAndServe()
}
