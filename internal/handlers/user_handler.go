package handlers

import (
	"encoding/json"
	"fin-tech-app/internal/model"
	"fin-tech-app/internal/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	err = h.UserService.CreateUser(r.Context(), user)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	response := map[string]string{
		"message": "User created successfully",
	}
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	user, err := h.UserService.GetUserById(r.Context(), id)

	if err != nil {
		log.Printf("Error fetching user with ID %v: %v", id, err)
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "User retreived successfully",
		"user":    user,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.UserService.GetUsers(r.Context())
	if err != nil {
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "User retreived successfully",
		"users":   users,
	}
	json.NewEncoder(w).Encode(response)
}
