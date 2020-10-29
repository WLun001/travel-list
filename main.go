package main

import (
	"bytes"
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"log"
	"os"
	travellist "travel-list/travel-list"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("Running application in %v environment", os.Getenv("APP_ENV"))
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	if err := readConfig(); err != nil {
		return err
	}
	port := os.Getenv("PORT")
	dbURI := viper.Get("DATABASE_URI").(string)
	r, err := travellist.NewRepo(dbURI)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	service := travellist.NewService(r)

	app := fiber.New()
	if !travellist.IsProduction() {
		app.Use(logger.New())
		app.Use(cors.New())
	}

	app.Static("/web", "web/dist/web")
	app.Get("/web/*", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("web/dist/web/index.html")
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Redirect("/web")
	})

	travellist.Routes(app, service)
	return app.Listen(fmt.Sprintf(":%s", port))
}

func readConfig() error {
	if travellist.IsProduction() {
		secret, err := accessSecret(os.Getenv("APP_SECRET_MANAGER_RESOURCE"))
		if err != nil {
			return err
		}
		viper.SetConfigType("yaml")
		if err := viper.ReadConfig(bytes.NewBuffer(secret)); err != nil {
			return err
		}
	} else {
		viper.SetConfigFile(".env.yaml")
		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	}
	return nil
}

func accessSecret(name string) ([]byte, error) {

	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create secretmanager client: %v", err)
	}

	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}

	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to access secret version: %v", err)
	}

	return result.Payload.GetData(), nil
}
