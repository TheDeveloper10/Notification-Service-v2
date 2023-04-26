package controller

import "notification-service/internal/repository"

func NewTestController() *Test {
	return &Test{}
}

func NewTemplateController(templateRepo repository.ITemplate) *Template {
	return &Template{
		templateRepo: templateRepo,
	}
}

func NewNotificationController() *Notification {
	return &Notification{}
}
