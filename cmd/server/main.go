package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/minab/internship-backend/config"
	"github.com/minab/internship-backend/internal/api"
	"github.com/minab/internship-backend/internal/db"
	"github.com/minab/internship-backend/internal/repository"
	"github.com/minab/internship-backend/internal/service"
)

func main() {
	cfg := config.Load()

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL environment variable required")
	}

	dbConn, err := db.Connect(connStr)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	userRepo := repository.NewUserRepository(dbConn)
	userService := service.NewUserService(userRepo)
	userHandler := api.NewUserHandler(userService)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/users/", userHandler.GetUser)

	log.Printf("Server running on port %s\n", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, mux); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
