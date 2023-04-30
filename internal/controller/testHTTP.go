package controller

import "github.com/gofiber/fiber/v2"

type TestHTTP struct {
}

func (ctrl *TestHTTP) Get(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func (ctrl *TestHTTP) Post(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusCreated)
}
