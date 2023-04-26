package dto

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TemplateBulkFilter struct {
	Page    uint32
	PerPage uint32
}

func (tbf *TemplateBulkFilter) Fill(c *fiber.Ctx) error {
	tbf.Page = 0
	tbf.PerPage = 20

	pageStr := c.Query("page")
	if pageStr != "" {
		pageNum, err := strconv.ParseUint(pageStr, 10, 32)
		if err == nil {
			tbf.Page = uint32(pageNum)
		}
	}

	perPageStr := c.Query("perPage")
	if perPageStr != "" {
		perPageNum, err := strconv.ParseUint(perPageStr, 10, 32)
		if err == nil {
			tbf.PerPage = uint32(perPageNum)
		}
	}

	return nil
}
