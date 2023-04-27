package controller

import (
	"notification-service/internal/dto"
	"notification-service/internal/service"

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

	errs := n.notificationSvc.Send(&body)
	if errs == nil || len(errs) <= 0 {
		return c.SendStatus(fiber.StatusCreated)
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"errors": errs,
	})
}
