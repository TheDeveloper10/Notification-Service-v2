package service

import (
	"notification-service/internal/dto"
	"notification-service/internal/util"
)

type Client struct {
}

func (svc *Client) NewClient(permissions *dto.Permissions) (*dto.Client, util.StatusCode) {
	// TODO: receive permissions, return dto.Client

	return nil, util.StatusSuccess
}

func (svc *Client) IssueToken(client *dto.Client) (string, util.StatusCode) {
	// TODO: receive dto.Client, return string

	return "", util.StatusSuccess
}
