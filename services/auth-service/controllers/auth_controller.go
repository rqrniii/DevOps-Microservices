package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rqrniii/DevOps-Microservices/services/auth-service/models"
	"github.com/rqrniii/DevOps-Microservices/services/common/database"
	myjwt "github.com/rqrniii/DevOps-Microservices/services/common/jwt"
	"golang.org/x/crypto/bcrypt"
)

// In-memory user store
//var users = map[string]models.User{}

// Register endpoint
func Register(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(input.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	_, err = database.DB.Exec(
		"INSERT INTO users (email, password) VALUES ($1, $2)",
		input.Email,
		string(hashedPassword),
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user registered"})
}

// Login endpoint
func Login(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var hashedPassword string
	err := database.DB.QueryRow(
		"SELECT password FROM users WHERE email=$1",
		input.Email,
	).Scan(&hashedPassword)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(input.Password),
	); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := myjwt.GenerateToken(input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Me endpoint (protected)
func Me(c *gin.Context) {
	userEmail, exists := c.Get("userEmail")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"email": userEmail})
}
