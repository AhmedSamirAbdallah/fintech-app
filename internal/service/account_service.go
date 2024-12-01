package service

import (
	"context"
	"fin-tech-app/internal/model"
	"fin-tech-app/internal/repository"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type AccountService struct {
	AccountRepo *repository.AccountRepository
	UserRepo    *repository.UserRepository
}

func (s *AccountService) CreateAccount(ctx context.Context, account *model.Account) error {
	_, err := s.UserRepo.GetUserById(ctx, account.UserID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("user with ID %s does not exist", account.UserID)
		}
		return fmt.Errorf("failed to retrieve user with ID %s: %w", account.UserID, err)
	}
	return s.AccountRepo.CreateAccount(ctx, account)
}

func (s *AccountService) GetAccountById(ctx context.Context, id string) (*model.Account, error) {
	return s.AccountRepo.GetAccountById(ctx, id)
}

func (s *AccountService) GetAccounts(ctx context.Context) ([]model.Account, error) {
	return s.AccountRepo.GetAccounts(ctx)
}

func (s *AccountService) DeleteAccount(ctx context.Context, id string) error {
	return s.AccountRepo.DeleteAccount(ctx, id)
}

func (s *AccountService) GetAccountBalance(ctx context.Context, id string) (float64, error) {
	return s.AccountRepo.GetAccountBalance(ctx, id)
}

func (s *AccountService) UpdateAccount(ctx context.Context, id string, updatedAccount *model.Account) error {
	return s.AccountRepo.UpdateAccount(ctx, id, updatedAccount)
}
