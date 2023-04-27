package controller

import "notification-service/internal/service"

func NewTestController() *Test {
	return &Test{}
}

func NewTemplateController(templateSvc *service.Template) *Template {
	return &Template{
		templateSvc: templateSvc,
	}
}

func NewNotificationController(templateSvc *service.Template, notificationSvc *service.Notification) *Notification {
	return &Notification{
		templateSvc:     templateSvc,
		notificationSvc: notificationSvc,
	}
}
