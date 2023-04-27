package service

import "notification-service/internal/repository"

func NewTemplateService(templateRepo repository.ITemplate) *Template {
	return &Template{
		templateRepo: templateRepo,
	}
}

func NewNotificationService(templateSvc *Template) *Notification {
	return &Notification{
		templateSvc: templateSvc,
	}
}
