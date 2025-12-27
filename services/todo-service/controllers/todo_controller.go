package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rqrniii/DevOps-Microservices/services/common/database"
	"github.com/rqrniii/DevOps-Microservices/services/todo-service/models"
)

var todos = []models.Todo{}

func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func AddTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email := c.GetString("email")
	query := `INSERT INTO todos (task, completed, email) VALUES ($1, $2, $3) RETURNING id, task, completed, email, created_at`
	row := database.DB.QueryRow(query, todo.Task, false, email)

	var newTodo models.Todo
	if err := row.Scan(&newTodo.ID, &newTodo.Task, &newTodo.Completed, &newTodo.Email, &newTodo.CreatedAt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newTodo)
}
