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
	dbURI := viper.Get("DATABASE_URI").(string)
	r, err := travellist.NewRepo(dbURI)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	service := travellist.NewService(r)

	app := fiber.New()
	app.Use(logger.New())

	app.Static("/", "web/dist/web/")
	travellist.Routes(app, service)
	return app.Listen(port)
}
