package controller

import (
	"notification-service/internal/dto"
	"notification-service/internal/service"
	"notification-service/internal/util"

	"github.com/gofiber/fiber/v2"
)

type Template struct {
	templateSvc *service.Template
}

func (t *Template) GetBulk(c *fiber.Ctx) error {
	filter := dto.TemplateBulkFilter{}
	if err := filter.Fill(c); err != nil {
		return err
	}

	templates, status := t.templateSvc.GetBulkTemplates(&filter)
	if status != util.StatusSuccess {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get templates")
	}

	return c.Status(fiber.StatusOK).JSON(templates)
}

func (t *Template) Create(c *fiber.Ctx) error {
	body := dto.Template{}
	if err := c.BodyParser(&body); err != nil {
		return err
	} else if err := body.Validate(); err != nil {
		return err
	}

	id, status := t.templateSvc.CreateTemplate(&body)
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

	template, status := t.templateSvc.GetTemplateByID(tid.ID)
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

	status := t.templateSvc.UpdateTemplate(tid.ID, &body)
	if status == util.StatusNotFound {
		return fiber.NewError(fiber.StatusNotFound, "Template not found")
	} else if status != util.StatusSuccess {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to replace template")
	}

	return c.SendStatus(fiber.StatusOK)
}

func (t *Template) DeleteByID(c *fiber.Ctx) error {
	tid := dto.TemplateID{}
	if err := tid.Fill(c); err != nil {
		return err
	}

	status := t.templateSvc.DeleteTemplate(tid.ID)
	if status == util.StatusNotFound {
		return fiber.NewError(fiber.StatusNotFound, "Template not found")
	} else if status != util.StatusSuccess {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete template")
	}

	return c.SendStatus(fiber.StatusOK)
}
