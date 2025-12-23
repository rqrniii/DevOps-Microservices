package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	httpserver "github.com/rqrniii/DevOps-Microservices/services/ai-service/internal/http"
)

func main() {
	// Load environment variables from .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: .env file not found, relying on system environment variables")
	}

	r := httpserver.SetupRouter()

	// CORS config
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	r.Run(":" + port)
}
