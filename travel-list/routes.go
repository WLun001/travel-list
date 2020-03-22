package travellist

import (
	"github.com/gofiber/fiber"
)

func Routes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/health", func(c *fiber.Ctx) {
		_ = c.Status(200).
			JSON(map[string]interface{}{
				"health": "ok",
				"status": 200,
			})
	})

}
