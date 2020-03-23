package travellist

import (
	"context"
	"github.com/gofiber/fiber"
	"net/http"
	"time"
)

type appService struct {
	Repository Repository
}

type Service interface {
	getTravels(c *fiber.Ctx)
	getTravel(c *fiber.Ctx)
	createTravel(c *fiber.Ctx)
	updateTravel(c *fiber.Ctx)
	deleteTravel(c *fiber.Ctx)
}

func NewService(r Repository) Service {
	return &appService{Repository: r}
}

func (a appService) getTravels(c *fiber.Ctx) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	travels, err := a.Repository.findAll(ctx)
	response(travels, err, c)
}

func (a appService) getTravel(c *fiber.Ctx) {
	id := c.Params("id")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	travel, err := a.Repository.findOne(ctx, id)
	response(travel, err, c)
}

func (a appService) createTravel(c *fiber.Ctx) {
	panic("implement me")
}

func (a appService) updateTravel(c *fiber.Ctx) {
	panic("implement me")
}

func (a appService) deleteTravel(c *fiber.Ctx) {
	panic("implement me")
}

func response(data interface{}, err error, c *fiber.Ctx) {
	if err != nil {
		_ = c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": err.Error(),
		})
	} else {
		_ = c.Status(200).JSON(data)
	}
}
