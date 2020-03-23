package travellist

import (
	"github.com/gofiber/fiber"
	"net/http"
)

func Routes(app *fiber.App, service Service) {
	api := app.Group("/api")
	api.Get("/health", func(c *fiber.Ctx) {
		_ = c.Status(http.StatusOK).
			JSON(map[string]interface{}{
				"health": "ok",
				"status": http.StatusOK,
			})
	})

	api.Get("/travels", service.getTravels)
	api.Get("/travels/:id", service.getTravel)

}
