package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Static("/", "web/dist/web/")

	app.Get("/health", func(c *fiber.Ctx) {
		c.Send("health ok")
	})

	panic(app.Listen(3000))
}
