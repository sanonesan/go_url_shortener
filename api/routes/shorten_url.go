package routes

import (
	"net/http"

	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	myDB "url_shortener/database"
	"url_shortener/helpers"
	"url_shortener/services"
)

func ShortenURL(c *gin.Context) {
	// Get Original URL from request
	original_url := c.PostForm("url")

	if !valid.IsURL(original_url) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Cannot parse URL: Entered value is not an URL",
		})
		return
	}

	// check for the domain error
	// users may abuse the shortener by shorting the domain `localhost:port` itself
	// leading to a inifite loop, so don't accept the domain for shortening
	if !helpers.RemoveDomainError(original_url) {
		c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{
			"error": "haha... nice try",
		})
		return
	}

	// enforce https
	// all url will be converted to https before storing in database
	original_url = helpers.EnforceHTTPProtocol(original_url)

	// Get database pointer
	var db *myDB.Database = myDB.DB
	db.OpenConnection()

	// Add new url
	if _, exists := db.StorageU2SU[original_url]; !exists {
		db.StorageU2SU[original_url] = shorten(original_url)
		c.JSON(http.StatusOK, gin.H{
			"message":   "Shortened url created",
			"shortened": db.StorageU2SU[original_url],
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message":   "Shortened url exists",
			"shortened": db.StorageU2SU[original_url],
		})
	}

	db.CloseConnection()
}

func shorten(original_url string) string {
	shortened_url := services.GenerateShortURL(original_url)

	var db *myDB.Database = myDB.DB
	db.OpenConnection()

	url, exists := db.StorageSU2U[shortened_url]
	for {
		if !exists {
			db.StorageSU2U[shortened_url] = original_url
			db.CloseConnection()
			return shortened_url
		}
		if url == original_url {
			db.StorageSU2U[shortened_url] = original_url
			db.CloseConnection()
			return shortened_url
		}
		if url == "" {
			db.StorageSU2U[shortened_url] = original_url
			db.CloseConnection()
			return shortened_url
		}

		shortened_url = services.GenerateShortURL(shortened_url)
		url, exists = db.StorageSU2U[shortened_url]
	}
}
