package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var JWTSecret string

func LoadConfig() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment")
	}

	JWTSecret = os.Getenv("JWT_SECRET")
	if JWTSecret == "" {
		log.Fatal("JWT_SECRET environment variable required")
	}
}
