package middleware

import (
	"notification-service/internal/util"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Authorization(requiredPermissions util.PermissionsNumeric) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		permissions := claims["permissions"].(util.PermissionsNumeric)

		if permissions.HasPermission(requiredPermissions) {
			return c.Next()
		} else {
			return fiber.NewError(fiber.StatusForbidden, "Access denied")
		}
	}
}
