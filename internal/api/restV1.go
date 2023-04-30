package api

import (
	"notification-service/internal/api/middleware"
	"notification-service/internal/controller"
	"notification-service/internal/repository"
	"notification-service/internal/service"
	"notification-service/internal/util"

	"github.com/gofiber/fiber/v2"
)

func SetUpRESTV1(app *fiber.App) {
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
		testCtrl         = controller.NewTestController()
		templateCtrl     = controller.NewTemplateController(templateSvc)
		notificationCtrl = controller.NewNotificationController(notificationSvc)
		clientCtrl       = controller.NewClientController(clientSvc)
	)

	api := app.Group("/v1")

	api.Use(middleware.CORS)

	api.Get("/test", testCtrl.Get)
	api.Post("/test", testCtrl.Post)

	api.Post("/clients", middleware.Authentication, middleware.Authorization(util.PermissionManageClients), clientCtrl.New)
	api.Post("/clients/token", clientCtrl.IssueToken)

	api.Get("/notifications", middleware.Authentication, middleware.Authorization(util.PermissionReadNotifications), notificationCtrl.GetBulk)
	api.Post("/notifications", middleware.Authentication, middleware.Authorization(util.PermissionSendNotifications), notificationCtrl.Send)

	api.Get("/templates", middleware.Authentication, middleware.Authorization(util.PermissionReadTemplates), templateCtrl.GetBulk)
	api.Post("/templates", middleware.Authentication, middleware.Authorization(util.PermissionWriteTemplates), templateCtrl.Create)
	api.Get("/templates/:templateID", middleware.Authentication, middleware.Authorization(util.PermissionReadTemplates), templateCtrl.GetByID)
	api.Put("/templates/:templateID", middleware.Authentication, middleware.Authorization(util.PermissionWriteTemplates), templateCtrl.ReplaceByID)
	api.Delete("/templates/:templateID", middleware.Authentication, middleware.Authorization(util.PermissionWriteTemplates), templateCtrl.DeleteByID)
}
