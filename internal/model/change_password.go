package model

import "time"

type PasswordResetToken struct {
	Token     string
	UserID    string
	ExpiresAt time.Time
}
