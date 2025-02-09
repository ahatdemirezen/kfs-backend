package routes

import (
	"kfs-backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupThreatPlansRoutes(app *fiber.App) {
	threatPlans := app.Group("/api/threatplans")

	threatPlans.Post("/", handlers.CreateThreatPlan)
	threatPlans.Get("/:threatPlanId", handlers.GetThreatPlan)
	threatPlans.Get("/analysis/:analysisInfoId", handlers.GetThreatPlansByAnalysisInfo)
	threatPlans.Put("/:threatPlanId", handlers.UpdateThreatPlan)
	threatPlans.Delete("/:threatPlanId", handlers.DeleteThreatPlan)

	// Birden fazla ThreatPlan oluşturma
	threatPlans.Post("/bulk", handlers.CreateMultipleThreatPlans)
}
