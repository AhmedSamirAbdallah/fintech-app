package handlers

import (
	"encoding/json"
	"fin-tech-app/internal/model"
	"fin-tech-app/internal/service"
	"log"
	"net/http"
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
	json.NewEncoder(w).Encode(account)
}

func (h *AccountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := h.AccountService.GetAccounts(r.Context())
	if err != nil {
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accounts)

}

func (h *AccountHandler) GetAccountById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
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
	json.NewEncoder(w).Encode(account)

}

func (h *AccountHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	err := h.AccountService.DeleteAccount(r.Context(), id)
	if err != nil {
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(account)

}
