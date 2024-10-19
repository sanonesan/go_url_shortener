package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	myDB "url_shortener/database"
)

func ResolveShortURL(c *gin.Context) {
	short_url := c.Param("url")

	var db *myDB.Database = myDB.DB
	db.OpenConnection()

	if url, exists := db.StorageSU2U[short_url]; exists {
		db.CloseConnection()
		c.JSON(http.StatusFound, gin.H{
			"message": "Link found",
			"url":     url,
		})
		return
	} else {
		db.CloseConnection()
		c.JSON(http.StatusNotFound, gin.H{
			"message":   "Link not found",
			"short_url": short_url,
		})
		return
	}
}
