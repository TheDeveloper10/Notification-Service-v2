package controller

import "github.com/gofiber/fiber/v2"

type Test struct {
}

func (ctrl *Test) Get(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func (ctrl *Test) Post(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusCreated)
}
