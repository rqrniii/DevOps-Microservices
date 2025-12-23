package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rqrniii/DevOps-Microservices/services/todo-service/middleware"
)

type Todo struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	Email     string `json:"email"`
	Completed bool   `json:"completed"` // ✅ ADDED
}

var todos = []Todo{}

// Existing endpoints
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
		ID:        len(todos) + 1,
		Task:      input.Task,
		Email:     userEmail,
		Completed: false, // ✅ ADDED - new todos start as incomplete
	}

	todos = append(todos, todo)
	c.JSON(http.StatusCreated, todo)
}

// ✅ NEW FUNCTION - Toggle Todo Completion
func toggleTodo(c *gin.Context) {
	idParam := c.Param("id")

	// convert id to int
	id := 0
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid todo id"})
		return
	}

	userEmail := c.GetString("email")

	// Find and toggle the todo
	for i, t := range todos {
		if t.ID == id && t.Email == userEmail {
			todos[i].Completed = !todos[i].Completed
			c.JSON(http.StatusOK, todos[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
}

func deleteTodo(c *gin.Context) {
	idParam := c.Param("id")

	// convert id to int
	id := 0
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid todo id"})
		return
	}

	userEmail := c.GetString("email")

	for i, t := range todos {
		if t.ID == id && t.Email == userEmail {
			// remove item from slice
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "todo deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
}

func createAITasks(c *gin.Context) {
	var req struct {
		Tasks []string `json:"tasks"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userEmail := c.GetString("email")

	// DELETE previous AI tasks - REMOVED per best practices
	// Now AI tasks are treated like regular tasks

	// ADD new AI tasks
	for _, task := range req.Tasks {
		task = strings.TrimSpace(task)
		if task == "" {
			continue
		}

		todo := Todo{
			ID:        len(todos) + 1,
			Task:      task, // ✅ REMOVED "AI: " prefix per best practices
			Email:     userEmail,
			Completed: false, // ✅ ADDED
		}
		todos = append(todos, todo)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "AI tasks added", "count": len(req.Tasks)})
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
		todoRoutes.POST("/ai", createAITasks)
		todoRoutes.PUT("/:id/toggle", toggleTodo)
		todoRoutes.DELETE("/:id", deleteTodo)
	}
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy", "service": "todo"})
	})

	r.Run(":8081")
}
