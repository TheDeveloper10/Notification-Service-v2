package api

import (
	"notification-service/internal/api/middleware"
	"notification-service/internal/controller"
	"notification-service/internal/repository"
	"notification-service/internal/service"

	"github.com/gofiber/fiber/v2"
)

func SetUpRESTV1(app *fiber.App) {
	var (
		templateRepo           = repository.NewBasicTemplateRepository()
		notificationRepo       = repository.NewBasicNotificationRepository()
		notificationSenderRepo = repository.NewBasicNotificationSenderRepository()
	)

	var (
		templateSvc     = service.NewTemplateService(templateRepo)
		notificationSvc = service.NewNotificationService(templateSvc, notificationSenderRepo, notificationRepo)
	)

	var (
		testCtrl         = controller.NewTestController()
		templateCtrl     = controller.NewTemplateController(templateSvc)
		notificationCtrl = controller.NewNotificationController(notificationSvc)
		clientCtrl       = controller.NewClientController()
	)

	api := app.Group("/v1")

	api.Use(middleware.CORS)

	api.Get("/test", testCtrl.Get)
	api.Post("/test", testCtrl.Post)

	api.Post("/clients", clientCtrl.New)
	api.Post("/clients/token", clientCtrl.IssueToken)

	api.Get("/notifications", middleware.Auth, notificationCtrl.GetBulk)
	api.Post("/notifications", middleware.Auth, notificationCtrl.Send)

	api.Get("/templates", middleware.Auth, templateCtrl.GetBulk)
	api.Post("/templates", middleware.Auth, templateCtrl.Create)
	api.Get("/templates/:templateID", middleware.Auth, templateCtrl.GetByID)
	api.Put("/templates/:templateID", middleware.Auth, templateCtrl.ReplaceByID)
	api.Delete("/templates/:templateID", middleware.Auth, templateCtrl.DeleteByID)
}
