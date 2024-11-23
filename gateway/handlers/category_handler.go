package handlers

import (
	"github.com/ecommerce/clients/product/http/clients"
	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

func (h *CategoryHandler) GetCategories(c *fiber.Ctx) error {
	res, err := clients.ListCategory()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.JSON(res)
}
