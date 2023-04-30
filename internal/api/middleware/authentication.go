package middleware

import (
	"notification-service/internal/config"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func InitAuthenticationMiddleware() {
	Authentication = jwtware.New(jwtware.Config{
		SigningKey:   []byte(config.Master.Service.Auth.TokenSigningKey),
		ErrorHandler: AuthenticationErrorHandler,
	})
}

var Authentication fiber.Handler
