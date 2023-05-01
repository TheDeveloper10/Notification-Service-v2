package service

import (
	"notification-service/internal/dto"
	"notification-service/internal/repository"
	"sync"
)

func NewTemplateService(templateRepo repository.ITemplate) *Template {
	return &Template{
		templateRepo: templateRepo,

		cache:   map[uint64]*dto.CachedTemplate{},
		cacheMu: sync.RWMutex{},
	}
}

func NewNotificationService(templateSvc *Template, notificationSenderRepo repository.INotificationSender, notificationRepo repository.INotification) *Notification {
	return &Notification{
		templateSvc: templateSvc,

		notificationSenderRepo: notificationSenderRepo,
		notificationRepo:       notificationRepo,
	}
}

func NewClientService(clientRepo repository.IClient) *Client {
	return &Client{
		clientRepo: clientRepo,

		activeClients:   map[string]*dto.ActiveClient{},
		activeClientsMu: sync.RWMutex{},
	}
}
