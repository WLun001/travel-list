package travellist

import (
	"context"
	"errors"
	"github.com/gofiber/fiber"
	"log"
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
	response(travels, http.StatusOK, err, c)
}

func (a appService) getTravel(c *fiber.Ctx) {
	id := c.Params("id")
	if id == "" {
		response(nil, http.StatusUnprocessableEntity, errors.New("id is not defined"), c)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	travel, err := a.Repository.findOne(ctx, id)
	response(travel, http.StatusOK, err, c)
}

func (a appService) createTravel(c *fiber.Ctx) {
	var travel Travel
	if err := c.BodyParser(&travel); err != nil {
		response(travel, http.StatusUnprocessableEntity, err, c)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := a.Repository.insertOne(ctx, &travel)
	response(travel, http.StatusOK, err, c)
}

func (a appService) updateTravel(c *fiber.Ctx) {
	id := c.Params("id")
	log.Println(id)
	if id == "" {
		response(nil, http.StatusUnprocessableEntity, errors.New("id is not defined"), c)
	}
	var travel Travel
	if err := c.BodyParser(&travel); err != nil {
		response(travel, http.StatusUnprocessableEntity, err, c)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := a.Repository.updateOne(ctx, id, &travel)
	response(nil, http.StatusNoContent, err, c)
}

func (a appService) deleteTravel(c *fiber.Ctx) {
	id := c.Params("id")
	log.Println(id)
	if id == "" {
		response(nil, http.StatusUnprocessableEntity, errors.New("id is not defined"), c)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := a.Repository.deleteOne(ctx, id)
	response(nil, http.StatusNoContent, err, c)
}

func response(data interface{}, httpStatus int, err error, c *fiber.Ctx) {
	if err != nil {
		_ = c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": err.Error(),
		})
	} else {
		if data != nil {
			_ = c.Status(httpStatus).JSON(data)
		} else {
			c.Status(httpStatus)
		}
	}
}
