package service

import (
	"notification-service/internal/config"
	"notification-service/internal/dto"
	"notification-service/internal/repository"
	"notification-service/internal/util"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Client struct {
	clientRepo repository.IClient
}

func (svc *Client) NewClient(permissions *dto.Permissions) (*dto.Client, util.StatusCode) {
	// TODO: receive permissions, return dto.Client

	return nil, util.StatusSuccess
}

func (svc *Client) IssueToken(clientCredentials *dto.ClientCredentials) (string, util.StatusCode) {
	clientObj, status := svc.clientRepo.GetClient(clientCredentials)
	if status != util.StatusSuccess {
		return "", status
	}

	token, err := svc.newToken(clientObj.Permissions)
	if err != nil {
		util.Logger.Error().Msg(err.Error())
		return "", util.StatusInternal
	}

	return token, util.StatusSuccess
}

func (svc *Client) newToken(permissions util.PermissionsNumeric) (string, error) {
	claims := jwt.MapClaims{
		"permissions": permissions,
		"exp":         time.Now().Add(time.Second * time.Duration(config.Master.Service.Auth.TokenExpiryTime)).Unix(),
	}

	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenObj.SignedString([]byte(config.Master.Service.Auth.TokenSigningKey))
	if err != nil {
		return "", err
	}

	return token, nil
}
