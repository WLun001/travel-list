package travellist

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"time"
)

type appService struct {
	Repository Repository
}

type Service interface {
	getTravels(c *fiber.Ctx) error
	getTravel(c *fiber.Ctx) error
	createTravel(c *fiber.Ctx) error
	updateTravel(c *fiber.Ctx) error
	deleteTravel(c *fiber.Ctx) error
}

func NewService(r Repository) Service {
	return &appService{Repository: r}
}

func (a *appService) getTravels(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	travels, err := a.Repository.findAll(ctx)
	return response(travels, http.StatusOK, err, c)
}

func (a *appService) getTravel(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return response(nil, http.StatusUnprocessableEntity, errors.New("id is not defined"), c)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	travel, err := a.Repository.findOne(ctx, id)
	return response(travel, http.StatusOK, err, c)
}

func (a *appService) createTravel(c *fiber.Ctx) error {
	var travel Travel
	if err := c.BodyParser(&travel); err != nil {
		return response(travel, http.StatusUnprocessableEntity, err, c)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := a.Repository.insertOne(ctx, &travel)
	return response(travel, http.StatusOK, err, c)
}

func (a *appService) updateTravel(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Println(id)
	if id == "" {
		return response(nil, http.StatusUnprocessableEntity, errors.New("id is not defined"), c)
	}
	var travel Travel
	if err := c.BodyParser(&travel); err != nil {
		return response(travel, http.StatusUnprocessableEntity, err, c)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := a.Repository.updateOne(ctx, id, &travel)
	return response(nil, http.StatusNoContent, err, c)
}

func (a *appService) deleteTravel(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Println(id)
	if id == "" {
		return response(nil, http.StatusUnprocessableEntity, errors.New("id is not defined"), c)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := a.Repository.deleteOne(ctx, id)
	return response(nil, http.StatusNoContent, err, c)
}

func response(data interface{}, httpStatus int, err error, c *fiber.Ctx) error {
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": err.Error(),
		})
	} else {
		if data != nil {
			return c.Status(httpStatus).JSON(data)
		} else {
			c.Status(httpStatus)
			return nil
		}
	}
}
