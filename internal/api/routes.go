package api

import (
	"net/http"

	"github.com/minab/internship-backend/internal/service"
)

// RegisterPublicRoutes sets up public endpoints: login and register.
func RegisterPublicRoutes(mux *http.ServeMux, userService *service.UserService, passwordResetService *service.PasswordResetService) {
	authHandler := NewAuthHandler(userService)
	userHandler := NewUserHandler(userService)

	mux.HandleFunc("/api/v1/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		authHandler.Login(w, r)
	})

	mux.HandleFunc("/api/v1/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		userHandler.CreateUser(w, r)
	})

	passwordResetHandler := NewPasswordResetHandler(passwordResetService)
	mux.HandleFunc("/api/v1/forgot-password", passwordResetHandler.ForgotPassword)
	mux.HandleFunc("/api/v1/reset-password", passwordResetHandler.ResetPassword)
}

func RegisterProtectedRoutes(mux *http.ServeMux, userService *service.UserService) {
	userHandler := NewUserHandler(userService)

	// /api/v1/users - GET only (listing users, protected)
	mux.HandleFunc("/api/v1/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			userHandler.ListUsers(w, r)
			return
		}
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	// /api/v1/users/{id} - GET
	mux.HandleFunc("/api/v1/users/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			userHandler.GetUser(w, r)
			return
		}
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	// /api/v1/users/update/{id} - PUT or PATCH
	mux.HandleFunc("/api/v1/users/update/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut || r.Method == http.MethodPatch {
			userHandler.UpdateUser(w, r)
			return
		}
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})
}
