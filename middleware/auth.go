package middleware

import "github.com/gin-gonic/gin"

func AuthMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Read token from header
		// Validate JWT
		// Check role match
		// Store user info in context
	}
}
