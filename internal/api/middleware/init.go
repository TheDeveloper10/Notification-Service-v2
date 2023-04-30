package middleware

import "notification-service/internal/util"

func InitMiddlewares() {
	InitCORSMiddleware()
	InitAuthMiddleware()

	util.Logger.Info().Msg("Initialized all middlewares")
}
