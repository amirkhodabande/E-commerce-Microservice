package presentation

import (
	"github.com/ecommerce/product/infrastructure/persistence"
	"github.com/ecommerce/product/presentation/interfaces/http/handlers"
	"github.com/ecommerce/product/presentation/interfaces/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, container *persistence.Container) {
	productHandler := handlers.NewProductHandler(container)
	categoryHandler := handlers.NewCategoryHandler(container)

	api := app.Group("api/products", middlewares.RequestIdMiddleware)

	api.Get("/", middlewares.ValidateListProduct, productHandler.GetProducts)

	api.Get("/categories", categoryHandler.GetCategories)
}
