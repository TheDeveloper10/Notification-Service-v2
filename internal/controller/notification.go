package controller

import "github.com/gofiber/fiber/v2"

type Notification struct {
}

func (n *Notification) GetBulk(c *fiber.Ctx) error {
	return nil
}

func (n *Notification) Send(c *fiber.Ctx) error {
	return nil
}
