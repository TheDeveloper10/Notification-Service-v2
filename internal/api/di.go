package api

import (
	"notification-service/internal/controller"
	"notification-service/internal/repository"
	"notification-service/internal/service"
)

var (
	templateRepo           = repository.NewBasicTemplateRepository()
	notificationRepo       = repository.NewBasicNotificationRepository()
	notificationSenderRepo = repository.NewBasicNotificationSenderRepository()
	clientRepo             = repository.NewBasicClientRepository()
)

var (
	templateSvc     = service.NewTemplateService(templateRepo)
	notificationSvc = service.NewNotificationService(templateSvc, notificationSenderRepo, notificationRepo)
	clientSvc       = service.NewClientService(clientRepo)
)

var (
	testHTTPCtrl = controller.NewTestHTTPController()
	testRMQCtrl  = controller.NewTestRMQController()

	statsHTTPCtrl = controller.NewStatsHTTPController(clientSvc, templateSvc)

	templateHTTPCtrl = controller.NewTemplateHTTPController(templateSvc)
	templateRMQCtrl  = controller.NewTemplateRMQController(templateSvc)

	notificationHTTPCtrl = controller.NewNotificationHTTPController(notificationSvc)
	notificationRMQCtrl  = controller.NewNotificationRMQController(notificationSvc)

	clientHTTPCtrl = controller.NewClientHTTPController(clientSvc)
)
