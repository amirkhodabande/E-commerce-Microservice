package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func RequestIdMiddleware(c *fiber.Ctx) error {
	requestID := c.Get("X-Request-ID")
	if requestID == "" {
		requestID = uuid.New().String()
	}

	c.Set("X-Request-ID", requestID)

	c.Locals("requestID", requestID)

	return c.Next()
}
