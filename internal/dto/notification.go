package dto

import "github.com/gofiber/fiber/v2"

type Notification struct {
	AppID        string            `json:"appId"`
	TemplateID   uint64            `json:"templateId"`
	Title        string            `json:"title"`
	Placeholders map[string]string `json:"placeholders"`

	Targets []NotificationTarget `json:"targets"`
}

func (n *Notification) Validate() error {
	aidLen := len(n.AppID)
	if aidLen <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "App ID must be at least 1 character")
	} else if aidLen > 16 {
		return fiber.NewError(fiber.StatusBadRequest, "App ID must be at most 16 characters")
	}

	titleLen := len(n.Title)
	if titleLen <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Title must be at least 1 character")
	} else if titleLen > 128 {
		return fiber.NewError(fiber.StatusBadRequest, "Title must be at most 128 characters")
	}

	for _, target := range n.Targets {
		if err := target.Validate(); err != nil {
			return err
		}
	}

	return nil
}
