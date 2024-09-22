package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ecommerce/user/presentation"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var app = fiber.New(fiber.Config{
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		return ctx.Status(http.StatusInternalServerError).JSON(map[string]any{
			// move to custom response objects
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

	presentation.RegisterRoutes(app)

	if err := app.Listen(*address); err != nil {
		panic(err)
	}

	fmt.Printf("Listening to %v \n", address)
}
