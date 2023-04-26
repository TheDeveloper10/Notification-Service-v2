package dto

import (
	"notification-service/internal/config"

	"github.com/gofiber/fiber/v2"
)

type Template struct {
	ID uint64 `json:"id"`

	Body     TemplateBody `json:"body"`
	Language string       `json:"language"`
	Type     string       `json:"type"`
}

func (t *Template) Validate() error {
	if err := t.Body.Validate(); err != nil {
		return err
	}

	hasLanguage := config.Master.Service.AllowedLanguages.Has(t.Language)
	if !hasLanguage {
		return fiber.NewError(fiber.StatusBadRequest, "Unknown language")
	}

	typeLen := len(t.Type)
	if typeLen <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Type must be at least 1 character")
	} else if typeLen > 8 {
		return fiber.NewError(fiber.StatusBadRequest, "Type must be at most 8 characters")
	}

	return nil
}
