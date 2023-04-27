package middleware

import (
	"encoding/json"
	"notification-service/internal/util"

	"github.com/gofiber/fiber/v2"
)

func Error(c *fiber.Ctx, err error) error {
	if e, ok := err.(*fiber.Error); ok {
		if e.Code == fiber.StatusUnprocessableEntity {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": "Unprocessable Entity",
			})
		}

		if e.Message == "" {
			return c.SendStatus(e.Code)
		} else {
			return c.Status(e.Code).JSON(fiber.Map{
				"error": e.Message,
			})
		}
	} else if _, ok := err.(*json.SyntaxError); ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}

	util.Logger.Info().Msg(err.Error())
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": "Internal Server Error",
	})
}
