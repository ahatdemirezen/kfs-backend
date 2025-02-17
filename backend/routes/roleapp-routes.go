package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	roleapp := app.Group("/api/roleapp")

	roleapp.Post("/",
		middleware.AuthenticateMiddleware, 
		middleware.IsUserMiddleware,      
		handlers.RoleAppHandler,         
	)
}
