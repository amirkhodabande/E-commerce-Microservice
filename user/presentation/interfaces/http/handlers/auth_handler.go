package handlers

import (
	"github.com/ecommerce/user/application"
	"github.com/ecommerce/user/domain/entities"
	"github.com/ecommerce/user/presentation/interfaces/http/data_objects"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	*application.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		AuthService: application.NewAuthService(),
	}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	params := c.Context().UserValue("params").(*data_objects.RegisterUserParams)

	registerData := &entities.UserEntity{
		Username: params.Email,
		Email:    params.Email,
		Password: params.Password,
	}

	if err := h.AuthService.Register(registerData); err != nil {
		return err
	}

	return c.SendString("Register success!")
}
