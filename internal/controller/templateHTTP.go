package controller

import (
	"notification-service/internal/dto"
	"notification-service/internal/service"
	"notification-service/internal/util"

	"github.com/gofiber/fiber/v2"
)

type TemplateHTTP struct {
	templateSvc *service.Template
}

func (ctrl *TemplateHTTP) GetBulk(c *fiber.Ctx) error {
	filter := dto.TemplateBulkFilter{}
	if err := filter.Fill(c); err != nil {
		return err
	}

	templates, status := ctrl.templateSvc.GetBulkTemplates(&filter)
	if status != util.StatusSuccess {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get templates")
	}

	return c.Status(fiber.StatusOK).JSON(templates)
}

func (ctrl *TemplateHTTP) Create(c *fiber.Ctx) error {
	body := dto.Template{}
	if err := c.BodyParser(&body); err != nil {
		return err
	} else if err := body.Validate(); err != nil {
		return err
	}

	id, status := ctrl.templateSvc.CreateTemplate(&body)
	if status != util.StatusSuccess {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create template")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"templateId": id})
}

func (ctrl *TemplateHTTP) GetByID(c *fiber.Ctx) error {
	tid := dto.ID{}
	if err := tid.FillTemplate(c); err != nil {
		return err
	}

	template, status := ctrl.templateSvc.GetTemplateByID(tid.ID)
	if status == util.StatusNotFound {
		return fiber.NewError(fiber.StatusNotFound, "Template not found")
	} else if status != util.StatusSuccess {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get template")
	}

	return c.Status(fiber.StatusOK).JSON(template)
}

func (ctrl *TemplateHTTP) ReplaceByID(c *fiber.Ctx) error {
	tid := dto.ID{}
	if err := tid.FillTemplate(c); err != nil {
		return err
	}

	body := dto.Template{}
	if err := c.BodyParser(&body); err != nil {
		return err
	} else if err := body.Validate(); err != nil {
		return err
	}

	status := ctrl.templateSvc.UpdateTemplate(tid.ID, &body)
	if status == util.StatusNotFound {
		return fiber.NewError(fiber.StatusNotFound, "Template not found")
	} else if status != util.StatusSuccess {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to replace template")
	}

	return c.SendStatus(fiber.StatusOK)
}

func (ctrl *TemplateHTTP) DeleteByID(c *fiber.Ctx) error {
	tid := dto.ID{}
	if err := tid.FillTemplate(c); err != nil {
		return err
	}

	status := ctrl.templateSvc.DeleteTemplate(tid.ID)
	if status == util.StatusNotFound {
		return fiber.NewError(fiber.StatusNotFound, "Template not found")
	} else if status != util.StatusSuccess {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete template")
	}

	return c.SendStatus(fiber.StatusOK)
}
