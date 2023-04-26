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
		}
	}

	LastTemplateIDStr := c.Query("lastTemplateId")
	if LastTemplateIDStr != "" {
		LastTemplateIDNum, err := strconv.ParseUint(LastTemplateIDStr, 10, 32)
		if err == nil && LastTemplateIDNum > 0 {
			tbf.LastTemplateID = LastTemplateIDNum
		}
	}

	return nil
}
