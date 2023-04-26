package controller

import (
	"notification-service/internal/dto"

	"github.com/gofiber/fiber/v2"
)

type Template struct {
}

func (t *Template) GetBulk(c *fiber.Ctx) error {
	return nil
}

func (t *Template) Create(c *fiber.Ctx) error {
	body := dto.Template{}
	if err := c.BodyParser(&body); err != nil {
		return err
	} else if err := body.Validate(); err != nil {
		return err
	}

	return nil
}

func (t *Template) GetByID(c *fiber.Ctx) error {
	tid := dto.TemplateID{}
	if err := tid.Fill(c); err != nil {
		return err
	}

	return nil
}

func (t *Template) ReplaceByID(c *fiber.Ctx) error {
	tid := dto.TemplateID{}
	if err := tid.Fill(c); err != nil {
		return err
	}

	return nil
}

func (t *Template) DeleteByID(c *fiber.Ctx) error {
	tid := dto.TemplateID{}
	if err := tid.Fill(c); err != nil {
		return err
	}

	return nil
}
