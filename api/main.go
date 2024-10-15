package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	"url_shortener/main/routes"
)

// func setupRoutes(app *fiber.App) {
// 	app.Get("/:url", routes.ResolveURL)
// 	app.Post("/:url", routes.ShortenURL)
// }

func main() {
	fmt.Println("hello World")
	// err := godotenv.Load()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//
	// app := fiber.New()
	// app.Use(logger.New())
	//
	// setupRoutes(app)
	//
	// log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
