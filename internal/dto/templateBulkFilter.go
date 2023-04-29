package dto

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TemplateBulkFilter struct {
	PerPage        uint32
	LastTemplateID uint64
}

func (tbf *TemplateBulkFilter) Fill(c *fiber.Ctx) error {
	tbf.LastTemplateID = 0
	tbf.PerPage = 20

	perPageStr := c.Query("perPage")
	if perPageStr != "" {
		perPageNum, err := strconv.ParseUint(perPageStr, 10, 32)
		if err == nil {
			tbf.PerPage = uint32(perPageNum)
		} else {
			return fiber.NewError(fiber.StatusBadRequest, "Per Page must be a positive integer")
		}

		if perPageNum > 100 {
			tbf.PerPage = 100
		}
	}

	lastTemplateIDStr := c.Query("lastTemplateId")
	if lastTemplateIDStr != "" {
		lastTemplateIDNum, err := strconv.ParseUint(lastTemplateIDStr, 10, 32)
		if err == nil {
			tbf.LastTemplateID = lastTemplateIDNum
		} else {
			return fiber.NewError(fiber.StatusBadRequest, "Last Template ID must be a positive integer")
		}
	}

	return nil
}
