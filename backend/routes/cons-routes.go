package routes

import (
	"kfs-backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupConsRoutes(app *fiber.App) {
	cons := app.Group("/api/cons")

	cons.Post("/", handlers.CreateCon)
	cons.Get("/:conId", handlers.GetCon)
	cons.Get("/analysis/:analysisInfoId", handlers.GetConsByAnalysisInfo)
	cons.Put("/:conId", handlers.UpdateCon)
	cons.Delete("/:conId", handlers.DeleteCon)

	// Birden fazla Cons oluşturma
	cons.Post("/bulk", handlers.CreateMultipleCons)
}
