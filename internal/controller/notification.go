package controller

import (
	"notification-service/internal/service"

	"github.com/gofiber/fiber/v2"
)

type Notification struct {
	templateSvc     *service.Template
	notificationSvc *service.Notification
}

func (n *Notification) GetBulk(c *fiber.Ctx) error {
	return nil
}

func (n *Notification) Send(c *fiber.Ctx) error {
	return nil
}
