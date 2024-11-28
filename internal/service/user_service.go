package service

import (
	"context"
	"fin-tech-app/internal/model"
	"fin-tech-app/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (s *UserService) CreateUser(ctx context.Context, user model.User) error {
	return s.Repo.CreateUser(ctx, user)
}

func (s *UserService) GetUserById(ctx context.Context, id string) (*model.User, error) {
	return s.Repo.GetUserById(ctx, id)
}
