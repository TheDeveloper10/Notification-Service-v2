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

func NewNotificationController() *Notification {
	return &Notification{}
}
