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

	errs, status := n.notificationSvc.Send(&body)
	if status == util.StatusSuccess {
		return c.SendStatus(fiber.StatusCreated)
	} else if status == util.StatusNotFound {
		return fiber.NewError(fiber.StatusNotFound, errs[0].Error())
	} else if status == util.StatusInternal {
		return fiber.NewError(fiber.StatusInternalServerError, errs[0].Error())
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"errors": errs,
	})
}
