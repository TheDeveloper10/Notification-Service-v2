package dto

import (
	"notification-service/internal/util"

	"github.com/gofiber/fiber/v2"
)

type TemplateBody struct {
	Email *string `json:"email"`
	SMS   *string `json:"sms"`
	Push  *string `json:"push"`
}

func (tb *TemplateBody) Validate() error {
	if tb.Email == nil && tb.SMS == nil && tb.Push == nil {
		return fiber.NewError(fiber.StatusBadRequest, "You must provide an Email, SMS or Push body")
	}

	if tb.Email != nil {
		l := len(*tb.Email)
		if l > 2048 {
			return fiber.NewError(fiber.StatusBadRequest, "Email body must be at most 2048 characters")
		} else if l <= 0 {
			return fiber.NewError(fiber.StatusBadRequest, "Email body must be at least 1 character")
		}
	}

	if tb.SMS != nil {
		l := len(*tb.SMS)
		if l > 2048 {
			return fiber.NewError(fiber.StatusBadRequest, "SMS body must be at most 2048 characters")
		} else if l <= 0 {
			return fiber.NewError(fiber.StatusBadRequest, "SMS body must be at least 1 character")
		}
	}

	if tb.Push != nil {
		l := len(*tb.Push)
		if l > 2048 {
			return fiber.NewError(fiber.StatusBadRequest, "Push body must be at most 2048 characters")
		} else if l <= 0 {
			return fiber.NewError(fiber.StatusBadRequest, "Push body must be at least 1 character")
		}
	}

	return nil
}

func (tb *TemplateBody) Fill(values map[string]string) {
	if tb.Email != nil {
		emailTemplate := util.TemplateString(*tb.Email)
		emailFilledTemplate := emailTemplate.Fill(values)
		tb.Email = &emailFilledTemplate
	}

	if tb.SMS != nil {
		smsTemplate := util.TemplateString(*tb.SMS)
		smsFilledTemplate := smsTemplate.Fill(values)
		tb.SMS = &smsFilledTemplate
	}

	if tb.Push != nil {
		pushTemplate := util.TemplateString(*tb.Push)
		pushFilledTemplate := pushTemplate.Fill(values)
		tb.Push = &pushFilledTemplate
	}
}
