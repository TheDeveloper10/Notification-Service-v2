package service

import "notification-service/internal/repository"

func NewTemplateService(templateRepo repository.ITemplate) *Template {
	return &Template{
		templateRepo: templateRepo,
	}
}

func NewNotificationService(templateSvc *Template, notificationSenderRepo repository.INotificationSender, notificationRepo repository.INotification) *Notification {
	return &Notification{
		templateSvc: templateSvc,

		notificationSenderRepo: notificationSenderRepo,
		notificationRepo:       notificationRepo,
	}
}
