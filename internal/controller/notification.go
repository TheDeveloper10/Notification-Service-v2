package controller

import (
	"notification-service/internal/dto"
	"notification-service/internal/service"
	"notification-service/internal/util"

	"github.com/gofiber/fiber/v2"
)

type Notification struct {
	notificationSvc *service.Notification
}

func (n *Notification) GetBulk(c *fiber.Ctx) error {
	return nil
}

func (n *Notification) Send(c *fiber.Ctx) error {
	body := dto.NotificationRequest{}
	if err := c.BodyParser(&body); err != nil {
		return err
	} else if err := body.Validate(); err != nil {
		return err
	}

	failSpot, status := n.notificationSvc.Send(&body)
	if failSpot == 0 {
		return c.SendStatus(fiber.StatusCreated)
	} else if failSpot == 1 {
		if status == util.StatusNotFound {
			return fiber.NewError(fiber.StatusNotFound, "Template not found")
		} else {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to get template")
		}
	}

	return nil
}
