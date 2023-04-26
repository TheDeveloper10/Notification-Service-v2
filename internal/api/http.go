package api

import (
	"notification-service/internal/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewFiberApp() *fiber.App {
	return fiber.New(fiber.Config{
		ErrorHandler: middleware.Error,
	})
}
