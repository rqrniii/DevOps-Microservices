package controllers

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/rqrniii/DevOps-Microservices/services/todo-service/models"
)

var todos = make(map[int64]models.Todo)
var mu sync.Mutex
var idCounter int64 = 1

func nextID() int64 {
	mu.Lock()
	defer mu.Unlock()
	id := idCounter
	idCounter++
	return id
}

func CreateTodo(c *gin.Context) {
	userEmail := c.GetString("userEmail") // from middleware
	var input models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.ID = nextID()
	input.Owner = userEmail
	todos[input.ID] = input

	c.JSON(http.StatusOK, input)
}

func ListTodos(c *gin.Context) {
	userEmail := c.GetString("userEmail")
	mu.Lock()
	defer mu.Unlock()
	result := []models.Todo{}
	for _, t := range todos {
		if t.Owner == userEmail {
			result = append(result, t)
		}
	}
	c.JSON(http.StatusOK, result)
}

func GetTodo(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)
	mu.Lock()
	defer mu.Unlock()
	t, ok := todos[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, t)
}

func UpdateTodo(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)
	var input models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	defer mu.Unlock()
	t, ok := todos[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	t.Title = input.Title
	t.Done = input.Done
	todos[id] = t
	c.JSON(http.StatusOK, t)
}

func DeleteTodo(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)
	mu.Lock()
	defer mu.Unlock()
	_, ok := todos[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	delete(todos, id)
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
