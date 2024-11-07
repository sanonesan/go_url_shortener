package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Test function for gin framework
func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
