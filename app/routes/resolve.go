package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	db "url_shortener/database"
	"url_shortener/handlers"
)

// Return Long Url by short one
func ResolveShortURL(c *gin.Context) {
	short_url := c.Param("url")

	// Connect to database
	DB, err := db.InitConnection()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"message":   "Server error",
			"short_url": err,
		})
	}
	// Handler for database
	h := handlers.New(DB)

	if original_url, err := h.GetLongByShort(short_url); original_url != "" && err == nil {
		c.JSON(http.StatusFound, gin.H{
			"message": "Link found",
			"url":     original_url,
		})
	} else if original_url == "" && err == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message":   "Original link not found",
			"short_url": short_url,
		})
	} else if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"message":   "Server error",
			"short_url": err,
		})
	}
}
