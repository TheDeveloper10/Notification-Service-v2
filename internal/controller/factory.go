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

func NewNotificationController(notificationSvc *service.Notification) *Notification {
	return &Notification{
		notificationSvc: notificationSvc,
	}
}

func NewClientController(clientSvc *service.Client) *Client {
	return &Client{
		clientSvc: clientSvc,
	}
}
