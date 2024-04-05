package user

import (
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct{}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	log.Println("received request to create a user")
	w.Write([]byte("User created!"))
}

func (h *Handler) FindByID(w http.ResponseWriter, r *http.Request) {
	log.Println("handling READ request - Method:", r.Method)
	user, exists := loadUsers()[r.PathValue("id")]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling UPDATE request - Method:", r.Method)
}

func (h *Handler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	log.Println("received DELETE request for user")
}

func (h *Handler) PatchByID(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
}
