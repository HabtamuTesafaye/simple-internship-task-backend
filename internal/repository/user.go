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

// GetUserByID retrieves a user by their ID from the database.
func (r *UserRepository) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	user := &model.User{}
	err := r.db.QueryRowContext(ctx, "SELECT id, full_name, email, phone_number, role, created_at FROM users WHERE id=$1", id).
		Scan(&user.ID, &user.FullName, &user.Email, &user.PhoneNumber, &user.Role, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// CreateUser inserts a new user into the database and returns the created user with its ID and creation timestamp.
func (r *UserRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	err := r.db.QueryRowContext(ctx,
		"INSERT INTO users (full_name, email, password, phone_number, role) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at",
		user.FullName, user.Email, user.Password, user.PhoneNumber, user.Role,
	).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser updates an existing user in the database and returns the updated user.
func (r *UserRepository) UpdateUser(ctx context.Context, id string, user *model.User) (*model.User, error) {
	err := r.db.QueryRowContext(ctx,
		"UPDATE users SET full_name=$1, email=$2, password=$3, phone_number=$4, role=$5 WHERE id=$6 RETURNING id, full_name, email, password, phone_number, role, created_at",
		user.FullName, user.Email, user.Password, user.PhoneNumber, user.Role, id,
	).Scan(&user.ID, &user.FullName, &user.Email, &user.Password, &user.PhoneNumber, &user.Role, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// ListUsers retrieves all users from the database.
func (r *UserRepository) ListUsers(ctx context.Context) ([]*model.User, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, full_name, email, phone_number, role, created_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		user := &model.User{}
		if err := rows.Scan(&user.ID, &user.FullName, &user.Email, &user.PhoneNumber, &user.Role, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// GetUserByEmail retrieves a user by their email from the database.
func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user := &model.User{}
	err := r.db.QueryRowContext(ctx, "SELECT id, full_name, email, password, phone_number, role, created_at FROM users WHERE email=$1", email).
		Scan(&user.ID, &user.FullName, &user.Email, &user.Password, &user.PhoneNumber, &user.Role, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}
