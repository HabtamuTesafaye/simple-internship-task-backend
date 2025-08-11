package service

import (
	"context"

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

func (s *UserService) GetUser(ctx context.Context, id string) (*model.User, error) {
	return s.repo.GetUserByID(ctx, id)
}

func (s *UserService) CreateUser(ctx context.Context, req *model.CreateUserRequest) (*model.User, error) {
	hashed, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		FullName:    req.FullName,
		Email:       req.Email,
		Password:    hashed,
		PhoneNumber: req.PhoneNumber,
		Role:        req.Role,
	}
	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) UpdateUser(ctx context.Context, id string, user *model.User) (*model.User, error) {
	return s.repo.UpdateUser(ctx, id, user)
}

func (s *UserService) ListUsers(ctx context.Context) ([]*model.User, error) {
	return s.repo.ListUsers(ctx)
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.repo.GetUserByEmail(ctx, email)
}
