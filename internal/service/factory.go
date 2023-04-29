package service

import (
	"notification-service/internal/dto"
	"notification-service/internal/repository"
	"sync"
)

func NewTemplateService(templateRepo repository.ITemplate) *Template {
	return &Template{
		templateRepo: templateRepo,
		cache:        map[uint64]*dto.Template{},
		cacheMu:      sync.RWMutex{},
	}
}

func NewNotificationService(templateSvc *Template, notificationSenderRepo repository.INotificationSender, notificationRepo repository.INotification) *Notification {
	return &Notification{
		templateSvc: templateSvc,

		notificationSenderRepo: notificationSenderRepo,
		notificationRepo:       notificationRepo,
	}
}
