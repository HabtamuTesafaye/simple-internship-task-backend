package repository

import (
	"context"
	"database/sql"

	"github.com/minab/internship-backend/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	user := &model.User{}
	err := r.db.QueryRowContext(ctx, "SELECT id, full_name, email, role FROM users WHERE id=$1", id).
		Scan(&user.ID, &user.FullName, &user.Email, &user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}
