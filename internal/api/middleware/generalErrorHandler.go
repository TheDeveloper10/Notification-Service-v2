package middleware

import (
	"encoding/json"
	"notification-service/internal/dto"
	"notification-service/internal/util"

	"github.com/gofiber/fiber/v2"
)

func GeneralErrorHandler(c *fiber.Ctx, err error) error {
	if e, ok := err.(*fiber.Error); ok {
		if e.Code == fiber.StatusUnprocessableEntity {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(&dto.Error{Error: "Unprocessable Entity"})
		}

		if e.Message == "" {
			return c.SendStatus(e.Code)
		} else {
			return c.Status(e.Code).JSON(&dto.Error{Error: e.Message})
		}
	} else if _, ok := err.(*json.SyntaxError); ok {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.Error{Error: "Invalid JSON format"})
	}

	util.Logger.Info().Msg(err.Error())
	return c.Status(fiber.StatusInternalServerError).JSON(&dto.Error{Error: "Internal Server Error"})
}
