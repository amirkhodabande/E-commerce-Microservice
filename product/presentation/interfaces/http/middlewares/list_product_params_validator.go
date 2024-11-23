package middlewares

import (
	"github.com/ecommerce/product/presentation/interfaces/http/data_objects"
	"github.com/gofiber/fiber/v2"
)

func ValidateListProduct(c *fiber.Ctx) error {
	params := new(data_objects.ListProductParams)
	c.QueryParser(params)

	customValidator := &customValidator{
		validator: validate,
	}

	if errs := customValidator.validateAndFormatErrors(params); errs != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(map[string]*map[string]string{
			"errors": errs,
		})
	}

	c.Locals("query-params", params)

	return c.Next()
}