package routes

import (
	"github.com/ecommerce/gateway/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterProductRoutes(app *fiber.App) {
	api := app.Group("api")

	productHandler := handlers.NewProductHandler()
	categoryHandler := handlers.NewCategoryHandler()

	api.Get("/products", productHandler.GetProducts)

	api.Get("/categories", categoryHandler.GetCategories)
}
