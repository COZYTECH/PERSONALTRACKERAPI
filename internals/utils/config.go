package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	DBUser   string
	DBPass   string
	DBName   string
	DBHost   string
	DBPort   string
}

// LoadConfig loads environment variables from .env
func LoadConfig() *Config {
	// Load .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := &Config{
		Port:   os.Getenv("PORT"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
		DBName: os.Getenv("DB_NAME"),
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
	}

	// Fail immediately if any required env is missing
	if cfg.Port == "" {
		log.Fatal("PORT not set in .env")
	}
	if cfg.DBUser == "" {
		log.Fatal("DB_USER not set in .env")
	}
	if cfg.DBPass == "" {
		log.Fatal("DB_PASS not set in .env")
	}
	if cfg.DBName == "" {
		log.Fatal("DB_NAME not set in .env")
	}
	if cfg.DBHost == "" {
		log.Fatal("DB_HOST not set in .env")
	}
	if cfg.DBPort == "" {
		log.Fatal("DB_PORT not set in .env")
	}

	return cfg
}
