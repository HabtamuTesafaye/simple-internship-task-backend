// cmd/server/main.go
package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/minab/internship-backend/config"
	"github.com/minab/internship-backend/internal/api"
	"github.com/minab/internship-backend/internal/middleware"
	"github.com/minab/internship-backend/internal/repository"
	"github.com/minab/internship-backend/internal/service"
)

// ...existing imports...

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	cfg := config.Load()

	userRepo := repository.NewUserRepository(cfg.Database)
	userService := service.NewUserService(userRepo)

	mux := http.NewServeMux()

	// Register only public routes here (e.g., login, register)
	api.RegisterPublicRoutes(mux, userService)

	// Register protected routes on a separate mux
	protectedMux := http.NewServeMux()
	api.RegisterProtectedRoutes(protectedMux, userService)

	// Protect all /api/v1/ routes except login/register
	mux.Handle("/api/v1/", middleware.JWTAuth(protectedMux))

	log.Printf("Server running on port %s\n", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, mux); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
