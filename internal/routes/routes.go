package routes

import (
	"net/http"

	"vibrox-core/internal/controller"

	"github.com/gin-gonic/gin"
)

// Register sets up the routes for the Gin router, including health checks and arena move handling.
func Register(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"service": "vibrox-core",
			"status":  "healthy",
			"role":    "systems-lab-gateway",
		})
	})
	router.POST("/arena/moves", controller.PlayArenaMove)
}
