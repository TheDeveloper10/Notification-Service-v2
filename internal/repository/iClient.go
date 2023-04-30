package repository

import (
	"notification-service/internal/dto"
	"notification-service/internal/repository/basic"
	"notification-service/internal/repository/mock"
	"notification-service/internal/util"
)

type IClient interface {
	CreateClient(clientObj *dto.Client) util.StatusCode
	GetClient(clientCredentials *dto.ClientCredentials) (*dto.ClientMetadata, util.StatusCode)
}

func NewBasicClientRepository() IClient {
	return &basic.ClientRepository{}
}

func NewMockClientRepository() IClient {
	return &mock.ClientRepository{}
}
