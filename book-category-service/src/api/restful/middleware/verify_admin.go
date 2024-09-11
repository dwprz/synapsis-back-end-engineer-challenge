package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func (m *Middleware) VerifyAdmin(c *fiber.Ctx) error {
	userData := c.Locals("user_data").(jwt.MapClaims)

	if role := userData["role"].(string); role != "ADMIN" {
		return c.Status(403).JSON(fiber.Map{"errors": "access denied"})
	}

	return c.Next()
}
