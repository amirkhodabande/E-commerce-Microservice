package presentation

import (
	"github.com/ecommerce/user/infrastructure/persistence"
	"github.com/ecommerce/user/presentation/interfaces/http/handlers"
	"github.com/ecommerce/user/presentation/interfaces/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, container *persistence.Container) {
	authHandler := handlers.NewAuthHandler(container)

	api := app.Group("api/user", middlewares.RequestIdMiddleware)

	api.Post("/register", middlewares.ValidateRegisterUser, authHandler.Register)
	api.Post("/login", middlewares.ValidateLoginUser, authHandler.Login)
}
