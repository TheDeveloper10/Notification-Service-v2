package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func InitCORSMiddleware() {
	CORS = cors.New()
}

var CORS fiber.Handler
