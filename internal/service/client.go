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
	idStr, err := util.GenerateString(16)
	if err != nil {
		return nil, util.StatusInternal
	}

	secretStr, err := util.GenerateString(128)
	if err != nil {
		return nil, util.StatusInternal
	}

	clientObj := dto.Client{
		ID:          idStr,
		Secret:      secretStr,
		Permissions: permissions.Permissions,
	}

	status := svc.clientRepo.CreateClient(&clientObj)
	if status == util.StatusSuccess {
		return &clientObj, status
	}

	return nil, status
}

func (svc *Client) IssueToken(clientCredentials *dto.ClientCredentials) (string, util.StatusCode) {
	var clientMetadata *dto.ClientMetadata
	if clientCredentials.ID != config.Master.Service.Auth.MasterClientID &&
		clientCredentials.Secret != config.Master.Service.Auth.MasterClientSecret {
		clientMetadataTmp, status := svc.clientRepo.GetClient(clientCredentials)
		if status != util.StatusSuccess {
			return "", status
		}

		clientMetadata = clientMetadataTmp
	} else {
		clientMetadata = &dto.ClientMetadata{
			ID:          clientCredentials.ID,
			Permissions: util.PermissionAll,
		}
	}

	token, err := svc.newToken(clientMetadata.Permissions)
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
