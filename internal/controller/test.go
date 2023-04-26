package controller

import "github.com/gofiber/fiber/v2"

type Test struct {
}

func (t *Test) Get(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func (t *Test) Post(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusCreated)
}
