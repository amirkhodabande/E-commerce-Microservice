package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ecommerce/user/infrastructure/persistence"
	"github.com/ecommerce/user/presentation"
	"github.com/ecommerce/user/presentation/interfaces/http/data_objects"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var app = fiber.New(fiber.Config{
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		status := http.StatusInternalServerError

		if customErr, ok := err.(*data_objects.ApiError); ok {
			status = customErr.Status
		}

		return ctx.Status(status).JSON(map[string]any{
			"message": err.Error(),
		})
	},
})

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	address := flag.String("serverPort", os.Getenv("APP_PORT"), "")
	flag.Parse()

	container := persistence.NewContainer()

	presentation.RegisterRoutes(app, container)

	if err := app.Listen(*address); err != nil {
		panic(err)
	}

	fmt.Printf("Listening to %v \n", address)
}
