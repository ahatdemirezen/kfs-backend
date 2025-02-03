package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func AdminAuthMiddleware(c *fiber.Ctx) error {
	return c.Next()
}