package controller

import (
	"notification-service/internal/dto"
	"notification-service/internal/service"
	"notification-service/internal/util"

	"github.com/gofiber/fiber/v2"
)

type ClientHTTP struct {
	clientSvc *service.Client
}

func (ctrl *ClientHTTP) New(c *fiber.Ctx) error {
	body := dto.Permissions{}
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	clientObj, status := ctrl.clientSvc.NewClient(&body)
	if status == util.StatusSuccess {
		return c.Status(fiber.StatusCreated).JSON(clientObj)
	} else {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create client")
	}
}

func (ctrl *ClientHTTP) IssueToken(c *fiber.Ctx) error {
	body := dto.ClientCredentials{}
	if err := c.BodyParser(&body); err != nil {
		return err
	} else if err := body.Validate(); err != nil {
		return err
	}

	token, status := ctrl.clientSvc.IssueToken(&body)
	if status == util.StatusSuccess {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"token": token,
		})
	} else if status == util.StatusError {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to issue token")
	}

	return fiber.NewError(fiber.StatusForbidden, "Invalid credentials")
}
