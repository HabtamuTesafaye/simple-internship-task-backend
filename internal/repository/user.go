package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/minab/internship-backend/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// GetUserByID retrieves a user by their ID from the database.
func (r *UserRepository) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	user := &model.User{}
	err := r.db.QueryRowContext(ctx, "SELECT id, full_name, email, role FROM users WHERE id=$1", id).
		Scan(&user.ID, &user.FullName, &user.Email, &user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// CreateUser inserts a new user into the database and returns the created user with its ID and creation timestamp.
func (r *UserRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	err := r.db.QueryRowContext(ctx,
		"INSERT INTO users (full_name, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id, created_at",
		user.FullName, user.Email, user.Password, user.Role,
	).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		fmt.Printf("Error inserting user: %v\n", err)
		return nil, err
	}
	return user, nil
}

// UpdateUser updates an existing user in the database and returns the updated user.
func (r *UserRepository) UpdateUser(ctx context.Context, id int, user *model.User) (*model.User, error) {
	err := r.db.QueryRowContext(ctx,
		"UPDATE users SET full_name=$1, email=$2, password=$3, role=$4 WHERE id=$5 RETURNING id, full_name, email, password, role, created_at",
		user.FullName, user.Email, user.Password, user.Role, id,
	).Scan(&user.ID, &user.FullName, &user.Email, &user.Password, &user.Role, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// ListUsers retrieves all users from the database.
func (r *UserRepository) ListUsers(ctx context.Context) ([]*model.User, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, full_name, email, role FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		user := &model.User{}
		if err := rows.Scan(&user.ID, &user.FullName, &user.Email, &user.Role); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
