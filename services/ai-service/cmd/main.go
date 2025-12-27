package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"

	"github.com/rqrniii/DevOps-Microservices/services/ai-service/internal/config"
	httpserver "github.com/rqrniii/DevOps-Microservices/services/ai-service/internal/http"
	commonjwt "github.com/rqrniii/DevOps-Microservices/services/common/jwt"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Println("Warning: .env file not found, using system env")
	}

	cfg := config.Load()

	commonjwt.LoadJWT()

	r := httpserver.SetupRouter()

	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
			"https://task-genius.app",
			"http://task-genius.app",
		},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Start server
	r.Run(":" + cfg.Port)
}
