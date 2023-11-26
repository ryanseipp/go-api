package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/abc/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("Hello, " + c.Params("name") + "!")
		}
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
