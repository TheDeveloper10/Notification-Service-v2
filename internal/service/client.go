package service

import (
	"notification-service/internal/config"
	"notification-service/internal/dto"
	"notification-service/internal/repository"
	"notification-service/internal/util"
	"sync"
	"time"
)

type Client struct {
	clientRepo repository.IClient

	activeClients   map[string]*dto.ActiveClient
	activeClientsMu sync.RWMutex
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

	if uint32(len(svc.activeClients)) >= config.Master.Service.Auth.MaxActiveClients {
		return "", util.StatusTooMany
	}

	token, err := util.GenerateString(128)
	if err != nil {
		return "", util.StatusInternal
	}

	svc.activeClientsMu.Lock()
	defer svc.activeClientsMu.Unlock()

	svc.activeClients[token] = &dto.ActiveClient{
		Metadata:     clientMetadata,
		InactiveTime: time.Now().Add(time.Second * time.Duration(config.Master.Service.Auth.TokenExpiryTime)).Unix(),
	}

	return token, util.StatusSuccess
}

func (svc *Client) GetActiveClientMetadataFromToken(token string) *dto.ActiveClient {
	svc.activeClientsMu.RLock()
	activeClient := svc.activeClients[token]
	svc.activeClientsMu.RUnlock()

	if activeClient != nil && activeClient.InactiveTime <= time.Now().Unix() {
		svc.activeClientsMu.Lock()
		delete(svc.activeClients, token)
		svc.activeClientsMu.Unlock()
		return nil
	}

	return activeClient
}

func (svc *Client) GetActiveClientsCount() int {
	return len(svc.activeClients)
}
