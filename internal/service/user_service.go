package service

import (
	"context"
	"fin-tech-app/internal/model"
	"fin-tech-app/internal/repository"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func (s *UserService) CreateUser(ctx context.Context, user model.User) error {
	return s.UserRepo.CreateUser(ctx, user)
}

func (s *UserService) GetUserById(ctx context.Context, id string) (*model.User, error) {
	return s.UserRepo.GetUserById(ctx, id)
}

func (s *UserService) GetUsers(ctx context.Context) ([]model.User, error) {
	return s.UserRepo.GetUsers(ctx)
}
