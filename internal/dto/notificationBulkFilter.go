package dto

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type NotificationBulkFilter struct {
	AppID              *string `json:"appId"`
	TemplateID         *uint64 `json:"templateId"`
	StartTime          *uint32 `json:"startTime"`
	EndTime            *uint32 `json:"endTime"`
	PerPage            uint32  `json:"perPage"`
	LastNotificationID uint64  `json:"lastNotificationId"`
}

func (nbf *NotificationBulkFilter) Validate() error {
	if nbf.TemplateID != nil && *nbf.TemplateID <= 0 {
		nbf.TemplateID = nil
	}

	if nbf.StartTime != nil && *nbf.StartTime <= 0 {
		nbf.StartTime = nil
	}

	if nbf.EndTime != nil && *nbf.EndTime <= 0 {
		nbf.EndTime = nil
	}

	if nbf.PerPage == 0 {
		nbf.PerPage = 20
	} else if nbf.PerPage > 100 {
		nbf.PerPage = 100
	}

	return nil
}

func (nbf *NotificationBulkFilter) Fill(c *fiber.Ctx) error {
	nbf.LastNotificationID = 0
	nbf.PerPage = 20

	appID := c.Query("appId")
	if appID != "" {
		nbf.AppID = &appID
	}

	templateIDStr := c.Query("templateId")
	if templateIDStr != "" {
		templateIDNum, err := strconv.ParseUint(templateIDStr, 10, 64)
		if err == nil && templateIDNum > 0 {
			nbf.TemplateID = &templateIDNum
		} else {
			return fiber.NewError(fiber.StatusBadRequest, "Template ID must be a positive integer")
		}
	}

	startTimeStr := c.Query("startTime")
	if startTimeStr != "" {
		startTimeNum, err := strconv.ParseUint(startTimeStr, 10, 32)
		if err == nil {
			t := uint32(startTimeNum)
			nbf.StartTime = &t
		} else {
			return fiber.NewError(fiber.StatusBadRequest, "Start Time must be a positive integer")
		}
	}

	endTimeStr := c.Query("endTime")
	if endTimeStr != "" {
		endTimeNum, err := strconv.ParseUint(endTimeStr, 10, 32)
		if err == nil {
			t := uint32(endTimeNum)
			nbf.EndTime = &t
		} else {
			return fiber.NewError(fiber.StatusBadRequest, "End Time must be a positive integer")
		}
	}

	perPageStr := c.Query("perPage")
	if perPageStr != "" {
		perPageNum, err := strconv.ParseUint(perPageStr, 10, 32)
		if err == nil {
			nbf.PerPage = uint32(perPageNum)
		} else {
			return fiber.NewError(fiber.StatusBadRequest, "Per Page must be a positive integer")
		}

		if perPageNum > 100 {
			nbf.PerPage = 100
		}
	}

	lastNotificationIDStr := c.Query("lastNotificationId")
	if lastNotificationIDStr != "" {
		lastNotificationIDNum, err := strconv.ParseUint(lastNotificationIDStr, 10, 32)
		if err == nil {
			nbf.LastNotificationID = lastNotificationIDNum
		} else {
			return fiber.NewError(fiber.StatusBadRequest, "Last Notification ID must be a positive integer")
		}
	}

	return nil
}
