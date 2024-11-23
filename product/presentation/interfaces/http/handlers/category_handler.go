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

type CategoryHandler struct {
	*application.CategoryService
}

func NewCategoryHandler(container *persistence.Container) *CategoryHandler {
	categoryRepository, err := container.Resolve(reflect.TypeOf((*contracts.CategoryRepository)(nil)))
	if err != nil {
		panic(err)
	}

	return &CategoryHandler{
		CategoryService: application.NewCategoryService(categoryRepository.(contracts.CategoryRepository)),
	}
}

func (h *CategoryHandler) GetCategories(ctx *fiber.Ctx) error {
	requestID := ctx.Locals("requestID").(string)
	fmt.Println("Request ID:", requestID)

	categories, err := h.CategoryService.GetCategories(requestID)
	if err != nil {
		return err
	}

	return ctx.JSON(data_objects.ListCategoryResponse{
		Success: true,
		Message: "Success",
		Data:    categories,
	})
}
