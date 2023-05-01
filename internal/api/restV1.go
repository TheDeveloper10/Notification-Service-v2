package api

import (
	"notification-service/internal/api/middleware"
	"notification-service/internal/util"

	"github.com/gofiber/fiber/v2"
)

func SetUpRESTV1(app *fiber.App) {
	api := app.Group("/v1")

	api.Use(middleware.CORS)

	api.Get("/test", testHTTPCtrl.Get)
	api.Post("/test", testHTTPCtrl.Post)

	api.Post("/clients", clientHTTPCtrl.Auth(util.PermissionManageClients), clientHTTPCtrl.New)
	api.Post("/clients/token", clientHTTPCtrl.IssueToken)

	api.Get("/notifications", clientHTTPCtrl.Auth(util.PermissionReadNotifications), notificationHTTPCtrl.GetBulk)
	api.Post("/notifications", clientHTTPCtrl.Auth(util.PermissionSendNotifications), notificationHTTPCtrl.Send)

	api.Get("/templates", clientHTTPCtrl.Auth(util.PermissionReadTemplates), templateHTTPCtrl.GetBulk)
	api.Post("/templates", clientHTTPCtrl.Auth(util.PermissionWriteTemplates), templateHTTPCtrl.Create)
	api.Get("/templates/:templateID", clientHTTPCtrl.Auth(util.PermissionReadTemplates), templateHTTPCtrl.GetByID)
	api.Put("/templates/:templateID", clientHTTPCtrl.Auth(util.PermissionWriteTemplates), templateHTTPCtrl.ReplaceByID)
	api.Delete("/templates/:templateID", clientHTTPCtrl.Auth(util.PermissionWriteTemplates), templateHTTPCtrl.DeleteByID)

	util.Logger.Info().Msg("Initialized REST V1 routes")
}
