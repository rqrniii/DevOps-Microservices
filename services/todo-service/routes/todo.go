package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rqrniii/DevOps-Microservices/services/todo-service/controllers"
	"github.com/rqrniii/DevOps-Microservices/services/todo-service/middleware"
)

func RegisterTodoRoutes(r *gin.Engine) {
	todos := r.Group("/todos")
	todos.Use(middleware.AuthMiddleware())
	{
		todos.GET("", controllers.GetTodos)
		todos.POST("", controllers.AddTodo)
	}
}
