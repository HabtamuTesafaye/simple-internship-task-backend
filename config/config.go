package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq" // or your DB driver
)

type Config struct {
	Port     string
	Database *sql.DB
}

func Load() *Config {
	port := getEnv("PORT", "8080")
	dbURL := getEnv("DATABASE_URL", "")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable required")
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}
	return &Config{
		Port:     port,
		Database: db,
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		log.Printf("Loaded environment variable %s=%s", key, value)
		return value
	}
	log.Printf("Environment variable %s not set, using fallback=%s", key, fallback)
	return fallback
}
