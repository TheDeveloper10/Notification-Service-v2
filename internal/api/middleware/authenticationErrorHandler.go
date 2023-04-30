package middleware

import (
	"notification-service/internal/dto"

	"github.com/gofiber/fiber/v2"
)

func AuthenticationErrorHandler(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(&dto.Error{Error: "Invalid or expired JWT"})
}
