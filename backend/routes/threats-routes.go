package routes

import (
	"kfs-backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupThreatsRoutes(app *fiber.App) {
	threats := app.Group("/api/threats")

	threats.Post("/", handlers.CreateThreat)
	threats.Get("/:threatId", handlers.GetThreat)
	threats.Get("/analysis/:analysisInfoId", handlers.GetThreatsByAnalysisInfo)
	threats.Put("/:threatId", handlers.UpdateThreat)
	threats.Delete("/:threatId", handlers.DeleteThreat)

	// Birden fazla Threat oluşturma
	threats.Post("/bulk", handlers.CreateMultipleThreats)
}
