package service

import (
	"context"
	"fmt"

	"github.com/minab/internship-backend/internal/model"
	"github.com/minab/internship-backend/internal/repository"
	"github.com/minab/internship-backend/internal/util"
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
	hashed, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashed
	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) UpdateUser(ctx context.Context, id int, user *model.User) (*model.User, error) {
	return s.repo.UpdateUser(ctx, id, user)
}

func (s *UserService) ListUsers(ctx context.Context) ([]*model.User, error) {
	return s.repo.ListUsers(ctx)
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.repo.GetUserByEmail(ctx, email)
}
