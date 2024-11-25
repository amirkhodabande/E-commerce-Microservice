package handlers

import (
	"fmt"

	"github.com/ecommerce/clients/product/http/clients"
	"github.com/ecommerce/clients/product/http/data_objects"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (h *ProductHandler) GetProducts(c *fiber.Ctx) error {
	requestID := c.Locals("requestID").(string)
	fmt.Println("GetProducts| requestID:", requestID)

	params := new(data_objects.ListProductData)
	c.QueryParser(params)

	errors := params.Validate()
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(map[string]any{"errors": errors})
	}

	res, err := clients.ListProduct(*params, requestID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.JSON(res)
}
