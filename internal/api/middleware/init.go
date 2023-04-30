package middleware

import "notification-service/internal/util"

func InitMiddlewares() {
	InitCORSMiddleware()
	InitAuthenticationMiddleware()

	util.Logger.Info().Msg("Initialized all middlewares")
}
