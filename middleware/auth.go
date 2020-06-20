package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "Bearer token2019" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you don't have the right!!"})
		c.Abort()
		return
	}
	c.Next()
}
