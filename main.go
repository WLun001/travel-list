package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	"github.com/spf13/viper"
	"log"
	travellist "travel-list/travel-list"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	viper.SetConfigFile(".env.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	port := viper.Get("PORT")

	app := fiber.New()
	app.Use(logger.New())

	app.Static("/", "web/dist/web/")
	travellist.Routes(app)
	return app.Listen(port)
}
