package api

import (
	"notification-service/internal/api/middleware"
	"notification-service/internal/util"

	"github.com/gofiber/fiber/v2"
)

func SetUpRESTV1(app *fiber.App) {
	api := app.Group("/v1")

	api.Use(middleware.CORS)

	api.Get("/test", statsHTTPCtrl.Middleware, testHTTPCtrl.Get)
	api.Post("/test", statsHTTPCtrl.Middleware, testHTTPCtrl.Post)

	// TODO: add auth middleware
	api.Get("/stats", statsHTTPCtrl.Middleware, statsHTTPCtrl.Get)

	api.Post("/clients", statsHTTPCtrl.Middleware, clientHTTPCtrl.Auth(util.PermissionManageClients), clientHTTPCtrl.New)
	api.Post("/clients/token", statsHTTPCtrl.Middleware, clientHTTPCtrl.IssueToken)

	api.Get("/notifications", statsHTTPCtrl.Middleware, clientHTTPCtrl.Auth(util.PermissionReadNotifications), notificationHTTPCtrl.GetBulk)
	api.Post("/notifications", statsHTTPCtrl.Middleware, clientHTTPCtrl.Auth(util.PermissionSendNotifications), notificationHTTPCtrl.Send)

	api.Get("/templates", statsHTTPCtrl.Middleware, clientHTTPCtrl.Auth(util.PermissionReadTemplates), templateHTTPCtrl.GetBulk)
	api.Post("/templates", statsHTTPCtrl.Middleware, clientHTTPCtrl.Auth(util.PermissionWriteTemplates), templateHTTPCtrl.Create)
	api.Get("/templates/:templateID", statsHTTPCtrl.Middleware, clientHTTPCtrl.Auth(util.PermissionReadTemplates), templateHTTPCtrl.GetByID)
	api.Put("/templates/:templateID", statsHTTPCtrl.Middleware, clientHTTPCtrl.Auth(util.PermissionWriteTemplates), templateHTTPCtrl.ReplaceByID)
	api.Delete("/templates/:templateID", statsHTTPCtrl.Middleware, clientHTTPCtrl.Auth(util.PermissionWriteTemplates), templateHTTPCtrl.DeleteByID)

	util.Logger.Info().Msg("Initialized REST V1 routes")
}
