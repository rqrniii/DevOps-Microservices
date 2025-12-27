package http

import (
	handler "github.com/rqrniii/DevOps-Microservices/services/ai-service/internal/http/handler/controllers"
	"github.com/rqrniii/DevOps-Microservices/services/ai-service/internal/http/middleware"
	"github.com/rqrniii/DevOps-Microservices/services/ai-service/internal/llm"
	"github.com/rqrniii/DevOps-Microservices/services/ai-service/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	llmClient := llm.NewClient()
	aiService := service.NewAIService(llmClient)
	aiHandler := handler.NewAIHandler(aiService)

	api := r.Group("/api/ai")
	api.Use(middleware.JWTAuth())
	{
		api.POST("/generate", aiHandler.Generate)
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	return r
}
