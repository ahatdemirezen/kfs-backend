package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupAdminRoutes(app *fiber.App) {
	admin := app.Group("/admin", middleware.AuthenticateMiddleware, middleware.SuperAdminAuthMiddleware)

	admin.Get("/users", handlers.GetAllUsers)       // Kullanıcıları listele
	admin.Delete("/users/:id", handlers.DeleteUser) // Kullanıcı sil
	admin.Patch("/users/:id/role", handlers.UpdateUserRole) // Kullanıcı rolünü güncelle
}
