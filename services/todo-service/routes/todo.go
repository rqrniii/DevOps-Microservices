package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rqrniii/DevOps-Microservices/services/todo-service/controllers"
	"github.com/rqrniii/DevOps-Microservices/services/todo-service/middleware"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	todoGroup := router.Group("/todos")
	todoGroup.Use(middleware.AuthMiddleware())
	{
		todoGroup.POST("/", controllers.CreateTodo)
		todoGroup.GET("/", controllers.ListTodos)
		todoGroup.GET("/:id", controllers.GetTodo)
		todoGroup.PUT("/:id", controllers.UpdateTodo)
		todoGroup.DELETE("/:id", controllers.DeleteTodo)
	}
}
