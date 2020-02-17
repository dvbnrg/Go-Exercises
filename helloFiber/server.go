package main

import "github.com/gofiber/fiber"

func main() {
	// Create new Fiber instance:
	app := fiber.New()

	// Create route on root path, "/":
	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
		// => "Hello, World!"
	})

	// Start server on "localhost" with port "80":
	app.Listen(80)
}
