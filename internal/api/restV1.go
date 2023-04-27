package api

import (
	"notification-service/internal/controller"
	"notification-service/internal/repository"
	"notification-service/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	)

	api := app.Group("/v1")

	api.Use(cors.New())

	api.Get("/test", testCtrl.Get)
	api.Post("/test", testCtrl.Post)

	api.Get("/notifications", notificationCtrl.GetBulk)
	api.Get("/notifications", notificationCtrl.Send)

	api.Get("/templates", templateCtrl.GetBulk)
	api.Post("/templates", templateCtrl.Create)
	api.Get("/templates/:templateID", templateCtrl.GetByID)
	api.Put("/templates/:templateID", templateCtrl.ReplaceByID)
	api.Delete("/templates/:templateID", templateCtrl.DeleteByID)
}
