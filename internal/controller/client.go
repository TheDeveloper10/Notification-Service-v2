package controller

import "github.com/gofiber/fiber/v2"

type Client struct {
}

func (ctrl *Client) New(c *fiber.Ctx) error {
	return nil
}

func (ctrl *Client) IssueToken(c *fiber.Ctx) error {
	return nil
}
