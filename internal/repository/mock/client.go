package mock

import (
	"notification-service/internal/dto"
	"notification-service/internal/util"
)

type ClientRepository struct {
}

func (cr *ClientRepository) CreateClient(clientObj *dto.Client) util.StatusCode {
	return util.StatusSuccess
}

func (cr *ClientRepository) GetClient(clientCredentials *dto.ClientCredentials) (*dto.ClientMetadata, util.StatusCode) {
	return nil, util.StatusSuccess
}
