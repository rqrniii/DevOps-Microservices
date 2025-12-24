package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rqrniii/DevOps-Microservices/services/auth-service/controllers"
	"github.com/rqrniii/DevOps-Microservices/services/auth-service/middleware"
)

func SetupRoutes(router *gin.Engine) {
	// Changed from /auth to /api/auth
	auth := router.Group("/api/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)

		auth.GET("/me", middleware.AuthMiddleware(), controllers.Me)
	}
}
