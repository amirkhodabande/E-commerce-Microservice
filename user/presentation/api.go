package presentation

import (
	"github.com/ecommerce/user/presentation/interfaces/http/handlers"
	"github.com/ecommerce/user/presentation/interfaces/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	authHandler := handlers.NewAuthHandler()

	api := app.Group("api/")

	api.Get("/register", middlewares.ValidateRegisterUser, authHandler.Register)
}
