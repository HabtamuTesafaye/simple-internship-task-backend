package service

import (
	"context"
	"fmt"

	"github.com/minab/internship-backend/internal/model"
	"github.com/minab/internship-backend/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(ctx context.Context, id int) (*model.User, error) {
	fmt.Printf("Fetching user with ID: %d\n", id)
	return s.repo.GetUserByID(ctx, id)
}

func (s *UserService) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) UpdateUser(ctx context.Context, id int, user *model.User) (*model.User, error) {
	return s.repo.UpdateUser(ctx, id, user)
}

func (s *UserService) ListUsers(ctx context.Context) ([]*model.User, error) {
	return s.repo.ListUsers(ctx)
}
