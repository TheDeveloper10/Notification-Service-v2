package dto

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TemplateID struct {
	ID uint64
}

func (tid *TemplateID) Fill(c *fiber.Ctx) error {
	str := c.Params("templateID")
	value, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Template ID must be a positive integer")
	}

	tid.ID = value
	return nil
}
