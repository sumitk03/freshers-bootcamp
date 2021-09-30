package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthouriseRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("admin") != "admin-pass" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "You have no right to access it! Try Again",
			})
			return
		}
		c.Next()
	}
}
