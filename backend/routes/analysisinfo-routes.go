package routes

import (
	"kfs-backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupAnalysisInfoRoutes(app *fiber.App) {
	analysis := app.Group("/api/analysis-info")

	analysis.Post("/", handlers.CreateAnalysisInfo)
	analysis.Get("/:analysisInfoId", handlers.GetAnalysisInfo)
	analysis.Put("/:analysisInfoId", handlers.UpdateAnalysisInfo)
	analysis.Delete("/:analysisInfoId", handlers.DeleteAnalysisInfo)
}
