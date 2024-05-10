package main

import (
	"log"
	"os"

	"github.com/Yash-sudo-web/urlshortnerredisgolang/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/shorten", routes.ShortenURL)
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	app.Use(logger.New())

	setupRoutes(app)

	app.Listen(os.Getenv("APP_PORT"))
}
