package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	todo.ID = len(todos) + 1
	todos = append(todos, todo)

	c.JSON(http.StatusCreated, todo)
}
