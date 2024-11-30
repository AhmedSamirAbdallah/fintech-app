package router

import (
	"fin-tech-app/internal/handlers"

	"github.com/gorilla/mux"
)

func RegisterAccountRoutes(mux *mux.Router, accountHandler *handlers.AccountHandler) {
	mux.HandleFunc("/api/accounts", accountHandler.CreateAccount).Methods("POST")
	mux.HandleFunc("/api/accounts/{id}", accountHandler.GetAccountById).Methods("GET")
	mux.HandleFunc("/api/accounts", accountHandler.GetAccounts).Methods("GET")
	mux.HandleFunc("/api/accounts/{id}/balance", accountHandler.GetAccountBalance).Methods("GET")
	mux.HandleFunc("/api/accounts/{id}", accountHandler.DeleteAccount).Methods("DELETE")
	mux.HandleFunc("/api/accounts/{id}", accountHandler.UpdateAccount).Methods("PUT")
}
