package dto

import (
	"github.com/gofiber/fiber/v2"
)

type TemplateBody struct {
	Email *string `json:"email"`
	SMS   *string `json:"sms"`
	Push  *string `json:"push"`
}

func (tb *TemplateBody) Validate() error {
	fields := 0

	if tb.Email != nil {
		fields++

		l := len(*tb.Email)
		if l > 2048 {
			return fiber.NewError(fiber.StatusBadRequest, "Email body must be at most 2048 characters")
		} else if l <= 0 {
			return fiber.NewError(fiber.StatusBadRequest, "Email body must be at least 1 character")
		}
	}

	if tb.SMS != nil {
		fields++

		l := len(*tb.SMS)
		if l > 2048 {
			return fiber.NewError(fiber.StatusBadRequest, "SMS body must be at most 2048 characters")
		} else if l <= 0 {
			return fiber.NewError(fiber.StatusBadRequest, "SMS body must be at least 1 character")
		}
	}

	if tb.Push != nil {
		fields++

		l := len(*tb.Push)
		if l > 2048 {
			return fiber.NewError(fiber.StatusBadRequest, "Push body must be at most 2048 characters")
		} else if l <= 0 {
			return fiber.NewError(fiber.StatusBadRequest, "Push body must be at least 1 character")
		}
	}

	if fields != 1 {
		return fiber.NewError(fiber.StatusBadRequest, "You must provide only one of 'email', 'sms', 'push'")
	}

	return nil
}
