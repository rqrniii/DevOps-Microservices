package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	commonjwt "github.com/rqrniii/DevOps-Microservices/services/common/jwt"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		token := strings.TrimPrefix(auth, "Bearer ")

		email, err := commonjwt.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		c.Set("email", email)
		c.Next()
	}
}
