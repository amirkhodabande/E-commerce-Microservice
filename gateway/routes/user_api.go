package routes

import (
	"github.com/ecommerce/gateway/handlers"
	"github.com/ecommerce/gateway/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App) {
	api := app.Group("api/user", middlewares.RequestIdMiddleware)

	authHandler := handlers.NewAuthHandler()

	api.Post("/register", authHandler.Register)
	api.Post("/login", authHandler.Login)
}
