package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rqrniii/DevOps-Microservices/services/todo-service/middleware"
)

type Todo struct {
	ID    int    `json:"id"`
	Task  string `json:"task"`
	Email string `json:"email"`
}

var todos = []Todo{}

func getTodos(c *gin.Context) {
	userEmail := c.GetString("email")
	userTodos := []Todo{}

	for _, t := range todos {
		if t.Email == userEmail {
			userTodos = append(userTodos, t)
		}
	}

	c.JSON(http.StatusOK, userTodos)
}

func createTodo(c *gin.Context) {
	var input struct {
		Task string `json:"task"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userEmail := c.GetString("email")

	todo := Todo{
		ID:    len(todos) + 1,
		Task:  input.Task,
		Email: userEmail,
	}

	todos = append(todos, todo)
	c.JSON(http.StatusCreated, todo)
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	todoRoutes := r.Group("/todos")
	todoRoutes.Use(middleware.JWTAuth())
	{
		todoRoutes.GET("", getTodos)
		todoRoutes.POST("", createTodo)
	}

	r.Run(":8081")
}
