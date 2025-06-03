package api

import (
	"net/http"

	"github.com/minab/internship-backend/internal/service"
)

func RegisterPublicRoutes(mux *http.ServeMux, userService *service.UserService) {
	authHandler := NewAuthHandler(userService)
	mux.HandleFunc("/api/v1/login", authHandler.Login)
}

func RegisterProtectedRoutes(mux *http.ServeMux, userService *service.UserService) {
	userHandler := NewUserHandler(userService)

	// /api/v1/users - GET and POST
	mux.HandleFunc("/api/v1/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			userHandler.CreateUser(w, r)
		case http.MethodGet:
			userHandler.ListUsers(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
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
