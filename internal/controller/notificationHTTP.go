package controller

import (
	"notification-service/internal/dto"
	"notification-service/internal/service"
	"notification-service/internal/util"

	"github.com/gofiber/fiber/v2"
)

type NotificationHTTP struct {
	notificationSvc *service.Notification
}

func (ctrl *NotificationHTTP) GetBulk(c *fiber.Ctx) error {
	filter := dto.NotificationBulkFilter{}
	if err := filter.Fill(c); err != nil {
		return err
	}

	notifications, status := ctrl.notificationSvc.GetBulkNotifications(&filter)
	if status != util.StatusSuccess {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get notifications")
	}

	return c.Status(fiber.StatusOK).JSON(notifications)
}

func (ctrl *NotificationHTTP) Send(c *fiber.Ctx) error {
	body := dto.NotificationRequest{}
	if err := c.BodyParser(&body); err != nil {
		return err
	} else if err := body.Validate(); err != nil {
		return err
	}

	errs, status := ctrl.notificationSvc.SendNotifications(&body)
	if status == util.StatusSuccess {
		return c.SendStatus(fiber.StatusCreated)
	} else if status == util.StatusNotFound {
		return fiber.NewError(fiber.StatusNotFound, errs[0])
	} else if status == util.StatusInternal {
		return fiber.NewError(fiber.StatusInternalServerError, errs[0])
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"errors": errs,
	})
}
