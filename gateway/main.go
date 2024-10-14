package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ecommerce/gateway/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var app = fiber.New(fiber.Config{
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		status := http.StatusInternalServerError

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

	routes.RegisterUserRoutes(app)

	if err := app.Listen(*address); err != nil {
		panic(err)
	}

	fmt.Printf("Listening to %v \n", address)

	// registerUser := data_objects.RegisterUserData{
	// 	Email:    "test@test.com",
	// 	Password: "test1234",
	// }

	// res, err := clients.RegisterUser(registerUser)

	// fmt.Printf("%+v", res)
	// fmt.Printf("%+v", err)

	// loginUser := data_objects.LoginUserData{
	// 	Email:    "test@test.com",
	// 	Password: "test1234",
	// }

	// res, err := clients.LoginUser(loginUser)

	// fmt.Printf("%+v", res)
	// fmt.Printf("%+v", err)
}
