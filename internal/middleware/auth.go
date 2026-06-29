package middleware

import (
	"net/http"

	"vibrox-core/internal/config"
	"vibrox-core/internal/logs"
	"vibrox-core/internal/proto/auth"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware handles authentication and authorization
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			logs.LogError(c, "Authorization header is missing")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		resp, err := config.AuthClient.ValidateToken(c, &auth.ValidateTokenRequest{Token: token})
		if err != nil {
			logs.LogError(c, "Failed to validate token: "+err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			c.Abort()
			return
		}

		if !resp.Valid {
			logs.LogError(c, "Invalid token: "+resp.Error)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
