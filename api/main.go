package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"

	"url_shortener/routes"
)

func setupRouter(r *gin.Engine) {
	r.POST("/short", routes.ShortenURL)
	r.GET("/:url", routes.ResolveShortURL)
}

func main() {
	logger := zap.Must(zap.NewDevelopment())
	defer logger.Sync()

	if err := godotenv.Load(); err != nil {
		logger.Fatal("Error loading .env file")
	}

	logger.Info("CREAT router...")
	r := gin.Default()
	logger.Info("Router CREATED!")

	logger.Info("SETUP router...")
	setupRouter(r)
	logger.Info("Router IS READY!")

	logger.Info("Running API...")
	r.Run(
		":" + os.Getenv("API_PORT"),
	)
}
