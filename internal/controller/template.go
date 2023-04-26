package controller

import (
	"notification-service/internal/dto"
	"notification-service/internal/repository"
	"notification-service/internal/util"

	"github.com/gofiber/fiber/v2"
)

type Template struct {
	templateRepo repository.ITemplate
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

	id, status := t.templateRepo.CreateTemplate(&body)
	if status != util.StatusSuccess {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create template")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"templateId": id})
}

func (t *Template) GetByID(c *fiber.Ctx) error {
	tid := dto.TemplateID{}
	if err := tid.Fill(c); err != nil {
		return err
	}

	template, status := t.templateRepo.GetTemplateByID(tid.ID)
	if status == util.StatusNotFound {
		return fiber.NewError(fiber.StatusNotFound, "Template not found")
	} else if status != util.StatusSuccess {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get template")
	}

	return c.Status(fiber.StatusOK).JSON(template)
}

func (t *Template) ReplaceByID(c *fiber.Ctx) error {
	tid := dto.TemplateID{}
	if err := tid.Fill(c); err != nil {
		return err
	}

	body := dto.Template{}
	if err := c.BodyParser(&body); err != nil {
		return err
	} else if err := body.Validate(); err != nil {
		return err
	}

	status := t.templateRepo.UpdateTemplate(tid.ID, &body)
	if status != util.StatusSuccess {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create template")
	}

	return c.SendStatus(fiber.StatusOK)
}

func (t *Template) DeleteByID(c *fiber.Ctx) error {
	tid := dto.TemplateID{}
	if err := tid.Fill(c); err != nil {
		return err
	}

	return nil
}
