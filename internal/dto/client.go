package dto

import "github.com/gofiber/fiber/v2"

type Client struct {
	ID     string `json:"id"`
	Secret string `json:"secret"`
}

func (c *Client) Validate() error {
	if len(c.ID) != 16 {
		return fiber.NewError(fiber.StatusBadRequest, "ID must be 16 characters long")
	}

	if len(c.Secret) != 128 {
		return fiber.NewError(fiber.StatusBadRequest, "Secret must be 128 characters long")
	}

	return nil
}
