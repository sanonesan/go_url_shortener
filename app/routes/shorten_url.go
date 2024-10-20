package routes

import (
	"net/http"
	"os"

	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	db "url_shortener/database"
	"url_shortener/handlers"
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

	// Add new url
	short_url, err := h.GetShortByLong(original_url)
	if short_url == "" && err == nil {
		short_url, err = shorten(original_url)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
				"message": "Server error",
				"error":   err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message":   "Shortened url created",
				"shortened": "http://" + os.Getenv("DOMAIN") + ":" + os.Getenv("API_PORT") + "/" + short_url,
			})
		}
	} else if short_url != "" && err == nil {
		c.JSON(http.StatusFound, gin.H{
			"message":   "Shortened url exists",
			"shortened": "http://" + os.Getenv("DOMAIN") + ":" + os.Getenv("API_PORT") + "/" + short_url,
		})
	} else if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"message": "Server error",
			"error":   err,
		})
	}
}

func shorten(original_url string) (string, error) {
	short_url := services.GenerateShortURL(original_url)

	// Connect to database
	DB, err := db.InitConnection()
	if err != nil {
		return "", err
	}
	// Handler for database
	h := handlers.New(DB)

	url, err := h.GetLongByShort(short_url)
	for {

		if url == "" && err == nil {
			h.AddURLPair(original_url, short_url)
			return short_url, nil
		}
		if url == original_url {
			return short_url, nil
		}
		if err != nil {
			return "", err
		}

		short_url = services.GenerateShortURL(short_url)
		url, err = h.GetLongByShort(short_url)
	}
}
