package main

import (
	"os"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello shit world")
		// => Hello shit world
	})
	app.Get("/greeting", func(c *fiber.Ctx) error {
		nameQueryParameter := c.Query("name")
		if nameQueryParameter != "" {
			return c.SendString("Hello " + nameQueryParameter + ", welcome to hell")
			// => Hello <name>, welcome to hell
		}
		return c.SendString("Who am i?")
	})

	app.Listen(":" + os.Getenv("PORT"))
}
