package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/minab/internship-backend/internal/model"
	"github.com/minab/internship-backend/internal/repository"
	"github.com/minab/internship-backend/internal/util"
)

type PasswordResetService struct {
	repo     *repository.PasswordResetRepository
	userRepo *repository.UserRepository
}

func NewPasswordResetService(repo *repository.PasswordResetRepository, userRepo *repository.UserRepository) *PasswordResetService {
	return &PasswordResetService{repo: repo, userRepo: userRepo}
}

func (s *PasswordResetService) GenerateToken(ctx context.Context, email string) (string, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	token := hex.EncodeToString(b)
	expiresAt := time.Now().Add(15 * time.Minute)
	if err := s.repo.CreateToken(ctx, token, user.ID, expiresAt); err != nil {
		return "", err
	}
	return token, nil
}

func (s *PasswordResetService) ResetPassword(ctx context.Context, token, newPassword string) error {
	t, err := s.repo.GetToken(ctx, token)
	if err != nil || t.ExpiresAt.Before(time.Now()) {
		return err
	}
	hashed, err := util.HashPassword(newPassword)
	if err != nil {
		return err
	}
	user := &model.User{ID: t.UserID, Password: hashed}
	if _, err := s.userRepo.UpdateUser(ctx, t.UserID, user); err != nil {
		return err
	}
	return s.repo.DeleteToken(ctx, token)
}
