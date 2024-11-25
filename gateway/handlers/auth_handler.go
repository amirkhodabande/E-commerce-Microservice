package handlers

import (
	"fmt"

	"github.com/ecommerce/clients/user/http/clients"
	"github.com/ecommerce/clients/user/http/data_objects"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	requestID := c.Locals("requestID").(string)
	fmt.Println("Register| requestID:", requestID)

	params := new(data_objects.RegisterUserData)
	c.BodyParser(&params)

	errors := params.Validate()
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(map[string]any{"errors": errors})
	}

	res, err := clients.RegisterUser(*params, requestID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.JSON(res)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	requestID := c.Locals("requestID").(string)
	fmt.Println("Login| requestID:", requestID)

	params := new(data_objects.LoginUserData)
	c.BodyParser(&params)

	errors := params.Validate()
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(map[string]any{"errors": errors})
	}

	res, err := clients.LoginUser(*params, requestID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.JSON(res)
}
