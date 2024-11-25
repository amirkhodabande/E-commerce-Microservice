package handlers

import (
	"fmt"
	"reflect"

	"github.com/ecommerce/user/application"
	"github.com/ecommerce/user/domain/contracts"
	"github.com/ecommerce/user/domain/entities"
	"github.com/ecommerce/user/infrastructure/persistence"
	"github.com/ecommerce/user/presentation/interfaces/http/data_objects"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	*application.AuthService
}

func NewAuthHandler(container *persistence.Container) *AuthHandler {
	userRepository, err := container.Resolve(reflect.TypeOf((*contracts.UserRepository)(nil)))
	if err != nil {
		panic(err)
	}

	return &AuthHandler{
		AuthService: application.NewAuthService(userRepository.(contracts.UserRepository)),
	}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	requestID := c.Locals("requestID").(string)
	fmt.Println("Register| Request ID:", requestID)

	params := c.Locals("params").(*data_objects.RegisterUserParams)

	user, err := entities.NewUserEntity(0, params.Email, params.Password)
	if err != nil {
		return err
	}

	token, err := h.AuthService.Register(user)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(data_objects.RegisterUserResponse{
		Success: true,
		Message: "Register success!",
		Token:   token,
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	requestID := c.Locals("requestID").(string)
	fmt.Println("Request ID:", requestID)

	params := c.Locals("params").(*data_objects.LoginUserParams)

	user, err := entities.NewUserEntity(0, params.Email, params.Password)
	if err != nil {
		return err
	}

	token, err := h.AuthService.Login(user)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(data_objects.LoginUserResponse{
		Success: true,
		Message: "Login success!",
		Token:   token,
	})
}
