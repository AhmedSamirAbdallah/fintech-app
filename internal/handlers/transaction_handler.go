package handlers

import (
	"encoding/json"
	"fin-tech-app/internal/model"
	"fin-tech-app/internal/service"
	"fmt"
	"log"
	"net/http"
)

type TransactionHandler struct {
	TransactionService *service.TransactionService
}

func (h *TransactionHandler) CreateDeposite(w http.ResponseWriter, r *http.Request) {
	var transaction model.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}
	err = h.TransactionService.CreateDeposite(r.Context(), &transaction)
	if err != nil {
		log.Printf("Error processing deposit: %v", err)
		http.Error(w, fmt.Sprintf("Error processing deposit: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message":     fmt.Sprintf("Deposit of %.2f successful for account %s", transaction.Amount, transaction.AccountID),
		"transaction": transaction,
	}
	json.NewEncoder(w).Encode(response)
}
func (h *TransactionHandler) CreateWithdraw(w http.ResponseWriter, r *http.Request) {
	var transaction model.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}
	err = h.TransactionService.CreateWithdraw(r.Context(), &transaction)
	if err != nil {
		log.Printf("Error processing withdraw: %v", err)
		http.Error(w, fmt.Sprintf("Error processing withdraw: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message":     fmt.Sprintf("Withdraw of %.2f successful for account %s", transaction.Amount, transaction.AccountID),
		"transaction": transaction,
	}
	json.NewEncoder(w).Encode(response)
}
