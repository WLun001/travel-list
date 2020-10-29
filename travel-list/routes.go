package travellist

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func Routes(app *fiber.App, service Service) {
	api := app.Group("/api")
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).
			JSON(map[string]interface{}{
				"health": "ok",
				"status": http.StatusOK,
			})
	})

	api.Get("/travels", service.getTravels)
	api.Get("/travels/:id", service.getTravel)
	api.Post("/travels", service.createTravel)
	api.Put("/travels/:id", service.updateTravel)
	api.Delete("/travels/:id", service.deleteTravel)

}
