package dto

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ID struct {
	ID uint64 `json:"id"`
}

func (id *ID) ValidateTemplate() error {
	if id.ID <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Template ID must be greater than 0")
	}

	return nil
}

func (id *ID) FillTemplate(c *fiber.Ctx) error {
	str := c.Params("templateID")
	value, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Template ID must be a positive integer")
	} else if value <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Template ID must be greater than 0")
	}

	id.ID = value
	return nil
}
