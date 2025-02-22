package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoleApplicationRoutes(app *fiber.App) {
	// Rol başvuru rotalarını grupla
	roleAppGroup := app.Group("/role-applications", middleware.AuthenticateMiddleware, middleware.AdminAuthMiddleware)

	// Başvuruları listeleme
	roleAppGroup.Get("/", handlers.AdminRoleApplicationsHandler)

	// Başvuru statüsünü güncelleme
	roleAppGroup.Patch("/:id", handlers.UpdateRoleApplicationStatusHandler)
}
