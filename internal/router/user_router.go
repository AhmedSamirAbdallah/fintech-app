package router

import (
	"fin-tech-app/internal/handlers"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(mux *mux.Router, userHandler *handlers.UserHandler) {

	mux.HandleFunc("/api/users", userHandler.CreateUser).Methods("POST")
	mux.HandleFunc("/api/users/{id}", userHandler.GetUserById).Methods("GET")

}
