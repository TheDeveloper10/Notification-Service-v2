package api

import (
	"notification-service/internal/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetUpRESTV1(app *fiber.App) {
	var (
		testCtrl         = controller.NewTestController()
		notificationCtrl = controller.NewNotificationController()
		templateCtrl     = controller.NewTemplateController()
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
