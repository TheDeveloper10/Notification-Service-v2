package controller

import (
	"notification-service/internal/dto"
	"notification-service/internal/service"
	"sync"
)

func NewTestHTTPController() *TestHTTP {
	return &TestHTTP{}
}

func NewTestRMQController() *TestRMQ {
	return &TestRMQ{}
}

func NewStatsHTTPController() *StatsHTTP {
	return &StatsHTTP{
		executionTimesHTTP:   map[string]*dto.ExecutionTimes{},
		executionTimesHTTPMu: sync.Mutex{},
	}
}

func NewTemplateHTTPController(templateSvc *service.Template) *TemplateHTTP {
	return &TemplateHTTP{
		templateSvc: templateSvc,
	}
}

func NewTemplateRMQController(templateSvc *service.Template) *TemplateRMQ {
	return &TemplateRMQ{
		templateSvc: templateSvc,
	}
}

func NewNotificationHTTPController(notificationSvc *service.Notification) *NotificationHTTP {
	return &NotificationHTTP{
		notificationSvc: notificationSvc,
	}
}

func NewNotificationRMQController(notificationSvc *service.Notification) *NotificationRMQ {
	return &NotificationRMQ{
		notificationSvc: notificationSvc,
	}
}

func NewClientHTTPController(clientSvc *service.Client) *ClientHTTP {
	return &ClientHTTP{
		clientSvc: clientSvc,
	}
}
