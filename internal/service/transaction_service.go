package service

import (
	"context"
	"fin-tech-app/internal/model"
	"fin-tech-app/internal/repository"
	"fmt"
	"log"
)

type TransactionService struct {
	TransactionRepository *repository.TransactionRepository
	AccountRepository     *repository.AccountRepository
}

func NewTransactionService(transactionRepo *repository.TransactionRepository, accountRepo *repository.AccountRepository) *TransactionService {
	return &TransactionService{
		TransactionRepository: transactionRepo,
		AccountRepository:     accountRepo,
	}
}

func (s *TransactionService) CreateDeposite(ctx context.Context, transaction *model.Transaction) error {
	balance, err := s.AccountRepository.GetAccountBalance(ctx, transaction.AccountID)
	if err != nil {
		return fmt.Errorf("error retrieving account balance: %v", err)
	}
	newBalance := transaction.Amount + balance
	account := &model.Account{
		ID:      transaction.ID,
		Balance: newBalance,
	}

	err = s.AccountRepository.UpdateAccount(ctx, transaction.AccountID, account)
	if err != nil {
		return fmt.Errorf("error while updating account balance for deposite: %v", err)
	}
	err = s.TransactionRepository.CreateTransaction(ctx, transaction)
	if err != nil {
		return fmt.Errorf("error creating transaction record: %v", err)
	}
	log.Printf("Deposit of %.2f completed successfully for account %s", transaction.Amount, transaction.AccountID)
	return nil
}

func (s *TransactionService) CreateWithdraw(ctx context.Context, transaction *model.Transaction) error {
	balance, err := s.AccountRepository.GetAccountBalance(ctx, transaction.AccountID)
	if err != nil {
		return fmt.Errorf("error retrieving account balance: %v", err)
	}
	if balance < transaction.Amount {
		return fmt.Errorf("insufficient funds: balance %.2f is less than withdrawal amount %.2f", balance, transaction.Amount)
	}
	newBalance := balance - transaction.Amount
	account := &model.Account{
		ID:      transaction.ID,
		Balance: newBalance,
	}

	err = s.AccountRepository.UpdateAccount(ctx, transaction.AccountID, account)
	if err != nil {
		return fmt.Errorf("error while updating account balance for withdrawal: %v", err)
	}
	err = s.TransactionRepository.CreateTransaction(ctx, transaction)
	if err != nil {
		return fmt.Errorf("error creating transaction record: %v", err)
	}
	log.Printf("Withdrawal of %.2f completed successfully for account %s", transaction.Amount, transaction.AccountID)
	return nil
}
