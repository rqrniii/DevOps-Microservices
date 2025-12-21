package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rqrniii/DevOps-Microservices/services/todo-service/routes"
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
		port = "8081"
	}

	log.Printf("Todo service listening on port %s", port)
	log.Println("JWT_SECRET:", os.Getenv("JWT_SECRET"))
	log.Printf("JWT_SECRET length: %d", len(os.Getenv("JWT_SECRET")))
	router.Run(":" + port)
}
