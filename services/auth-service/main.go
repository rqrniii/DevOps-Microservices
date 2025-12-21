package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rqrniii/DevOps-Microservices/services/auth-service/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	router := gin.Default()
	router.SetTrustedProxies([]string{"10.0.0.0/8"})
	routes.SetupRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Auth service port
	}

	log.Printf("Auth service listening on port %s", port)
	log.Printf("JWT_SECRET length: %d", len(os.Getenv("JWT_SECRET")))
	router.Run(":" + port)
}
