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

	api.Post("/clients", clientHTTPCtrl.Authentication, middleware.Authorization(util.PermissionManageClients), clientHTTPCtrl.New)
	api.Post("/clients/token", clientHTTPCtrl.IssueToken)

	api.Get("/notifications", clientHTTPCtrl.Authentication, middleware.Authorization(util.PermissionReadNotifications), notificationHTTPCtrl.GetBulk)
	api.Post("/notifications", clientHTTPCtrl.Authentication, middleware.Authorization(util.PermissionSendNotifications), notificationHTTPCtrl.Send)

	api.Get("/templates", clientHTTPCtrl.Authentication, middleware.Authorization(util.PermissionReadTemplates), templateHTTPCtrl.GetBulk)
	api.Post("/templates", clientHTTPCtrl.Authentication, middleware.Authorization(util.PermissionWriteTemplates), templateHTTPCtrl.Create)
	api.Get("/templates/:templateID", clientHTTPCtrl.Authentication, middleware.Authorization(util.PermissionReadTemplates), templateHTTPCtrl.GetByID)
	api.Put("/templates/:templateID", clientHTTPCtrl.Authentication, middleware.Authorization(util.PermissionWriteTemplates), templateHTTPCtrl.ReplaceByID)
	api.Delete("/templates/:templateID", clientHTTPCtrl.Authentication, middleware.Authorization(util.PermissionWriteTemplates), templateHTTPCtrl.DeleteByID)

	util.Logger.Info().Msg("Initialized REST V1 routes")
}
