package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func home(c *fiber.Ctx) error {
	return c.SendString("Hello shit world")
}

func greeting(c *fiber.Ctx) error {
	nameQueryParameter := c.Query("name")
	if nameQueryParameter != "" {
		return c.SendString("Hello " + nameQueryParameter + ", welcome to hell")
		// => Hello <name>, welcome to hell
	}
	return c.SendString("Who am i?")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", home)
	app.Get("/greeting", greeting)
}

func main() {
	port := os.Getenv("PORT")
	if os.Getenv("PORT") == "" {
		port = "3000"
	}

	app := fiber.New()
	setupRoutes(app)

	app.Listen(":" + port)
}
