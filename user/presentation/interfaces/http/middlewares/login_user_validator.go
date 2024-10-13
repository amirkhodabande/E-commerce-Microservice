package middlewares

import (
	"github.com/ecommerce/user/presentation/interfaces/http/data_objects"
	"github.com/gofiber/fiber/v2"
)

func ValidateLoginUser(c *fiber.Ctx) error {
	params := new(data_objects.LoginUserParams)
	c.BodyParser(&params)

	customValidator := &customValidator{
		validator: validate,
	}

	if errs := customValidator.validateAndFormatErrors(params); errs != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(map[string]*map[string]string{
			"errors": errs,
		})
	}

	c.Locals("params", params)

	return c.Next()
}
