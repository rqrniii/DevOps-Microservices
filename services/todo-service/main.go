package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rqrniii/DevOps-Microservices/services/common/config"
	"github.com/rqrniii/DevOps-Microservices/services/common/database"
	"github.com/rqrniii/DevOps-Microservices/services/common/jwt"
	"github.com/rqrniii/DevOps-Microservices/services/todo-service/middleware"
	"github.com/rqrniii/DevOps-Microservices/services/todo-service/models"
)

func getTodos(c *gin.Context) {
	email := c.GetString("email")
	rows, err := database.DB.Query("SELECT id, task, completed, email, created_at FROM todos WHERE email=$1", email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	todos := []models.Todo{}
	for rows.Next() {
		var t models.Todo
		if err := rows.Scan(&t.ID, &t.Task, &t.Completed, &t.Email, &t.CreatedAt); err != nil {
			fmt.Println("DB scan error:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		todos = append(todos, t)
	}

	c.JSON(http.StatusOK, todos)
}

func createTodo(c *gin.Context) {
	var input struct {
		Task string `json:"task"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email := c.GetString("email")

	query := `INSERT INTO todos (task, completed, email) VALUES ($1, $2, $3) 
			  RETURNING id, task, completed, email, created_at`
	row := database.DB.QueryRow(query, input.Task, false, email)

	var todo models.Todo
	if err := row.Scan(&todo.ID, &todo.Task, &todo.Completed, &todo.Email, &todo.CreatedAt); err != nil {
		fmt.Println("INSERT scan error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

func toggleTodo(c *gin.Context) {
	idParam := c.Param("id")
	id := 0
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid todo id"})
		return
	}

	email := c.GetString("email")

	var completed bool
	query := `UPDATE todos SET completed = NOT completed 
			  WHERE id=$1 AND email=$2 RETURNING completed`
	err = database.DB.QueryRow(query, id, email).Scan(&completed)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id, "completed": completed})
}

func deleteTodo(c *gin.Context) {
	idParam := c.Param("id")
	id := 0
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid todo id"})
		return
	}

	email := c.GetString("email")
	res, err := database.DB.Exec("DELETE FROM todos WHERE id=$1 AND email=$2", id, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	count, _ := res.RowsAffected()
	if count == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "todo deleted"})
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
	count := 0

	for _, task := range req.Tasks {
		task = strings.TrimSpace(task)
		if task == "" {
			continue
		}

		query := `INSERT INTO todos (task, completed, email) 
				  VALUES ($1, $2, $3)`
		_, err := database.DB.Exec(query, task, false, userEmail)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		count++
	}

	c.JSON(http.StatusCreated, gin.H{"message": "AI tasks added", "count": count})
}

func main() {
	config.LoadConfig()
	jwt.LoadJWT()

	database.Connect()

	r := gin.Default()

	// ✅ FIXED: Dynamic CORS based on environment
	allowedOrigins := []string{
		"http://localhost:5173",
		"https://task-genius.app",
		"http://task-genius.app",
	}

	// Check for custom allowed origins from environment
	if customOrigins := os.Getenv("ALLOWED_ORIGINS"); customOrigins != "" {
		allowedOrigins = strings.Split(customOrigins, ",")
	}

	r.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// ✅ Routes remain at /api/todos
	todoRoutes := r.Group("/api/todos")
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
