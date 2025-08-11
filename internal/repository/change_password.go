package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/minab/internship-backend/internal/model"
)

type PasswordResetRepository struct {
	db *sql.DB
}

func NewPasswordResetRepository(db *sql.DB) *PasswordResetRepository {
	return &PasswordResetRepository{db: db}
}

func (r *PasswordResetRepository) CreateToken(ctx context.Context, token, userID string, expiresAt time.Time) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO password_reset_tokens (token, user_id, expires_at) VALUES ($1, $2, $3)", token, userID, expiresAt)
	return err
}

func (r *PasswordResetRepository) GetToken(ctx context.Context, token string) (*model.PasswordResetToken, error) {
	row := r.db.QueryRowContext(ctx, "SELECT token, user_id, expires_at FROM password_reset_tokens WHERE token=$1", token)
	var t model.PasswordResetToken
	if err := row.Scan(&t.Token, &t.UserID, &t.ExpiresAt); err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *PasswordResetRepository) DeleteToken(ctx context.Context, token string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM password_reset_tokens WHERE token=$1", token)
	return err
}
