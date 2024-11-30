package router

import (
	"fin-tech-app/internal/handlers"

	"github.com/gorilla/mux"
)

func RegisterTranscationRoutes(mux *mux.Router, transactionHandler *handlers.TransactionHandler) {
	mux.HandleFunc("/api/accounts/deposit", transactionHandler.CreateDeposite).Methods("POST")
	mux.HandleFunc("/api/accounts/withdraw", transactionHandler.CreateWithdraw).Methods("POST")
}
