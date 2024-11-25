package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func RequestIdMiddleware(c *fiber.Ctx) error {
	requestID := uuid.New().String()

	c.Locals("requestID", requestID)

	return c.Next()
}
