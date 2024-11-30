package handlers

import (
	"encoding/json"
	"fin-tech-app/internal/model"
	"fin-tech-app/internal/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type AccountHandler struct {
	AccountService *service.AccountService
}

func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account model.Account

	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err = h.AccountService.CreateAccount(r.Context(), &account)
	if err != nil {
		log.Printf("Error creating account: %v", err)
		http.Error(w, "Failed to create account", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]string{
		"message": "Account created successfully",
	}
	json.NewEncoder(w).Encode(response)
}

func (h *AccountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := h.AccountService.GetAccounts(r.Context())
	if err != nil {
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message":  "Accounts retrieved successfully",
		"accounts": accounts,
	}
	json.NewEncoder(w).Encode(response)

}

func (h *AccountHandler) GetAccountById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	account, err := h.AccountService.GetAccountById(r.Context(), id)
	if err != nil {
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Account retrieved successfully",
		"account": account,
	}
	json.NewEncoder(w).Encode(response)

}

func (h *AccountHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := h.AccountService.DeleteAccount(r.Context(), id)
	if err != nil {
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{
		"message": "Account deleted successfully",
		"id":      id,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *AccountHandler) GetAccountBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	balance, err := h.AccountService.GetAccountBalance(r.Context(), id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to retrieve balance: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Balance retrieved successfully",
		"balance": balance,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *AccountHandler) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var account model.Account
	json.NewDecoder(r.Body).Decode(&account)
	err := h.AccountService.UpdateAccount(r.Context(), id, &account)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update account: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Account updated successfully",
	}
	json.NewEncoder(w).Encode(response)
}
