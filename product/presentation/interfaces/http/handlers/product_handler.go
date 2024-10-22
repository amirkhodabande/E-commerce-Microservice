package handlers

import (
	"fmt"
	"reflect"

	"github.com/ecommerce/product/application"
	"github.com/ecommerce/product/domain/contracts"
	"github.com/ecommerce/product/infrastructure/persistence"
	"github.com/ecommerce/product/presentation/interfaces/http/data_objects"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	*application.ProductService
}

func NewProductHandler(container *persistence.Container) *ProductHandler {
	productRepository, err := container.Resolve(reflect.TypeOf((*contracts.ProductRepository)(nil)))
	if err != nil {
		panic(err)
	}

	return &ProductHandler{
		ProductService: application.NewProductService(productRepository.(contracts.ProductRepository)),
	}
}

func (h *ProductHandler) GetProducts(ctx *fiber.Ctx) error {
	requestID := ctx.Locals("requestID").(string)
	fmt.Println("Request ID:", requestID)

	products, err := h.ProductService.GetProducts(requestID)
	if err != nil {
		return err
	}

	return ctx.JSON(data_objects.ListProductResponse{
		Success: true,
		Message: "Success",
		Data:    products,
	})
}
