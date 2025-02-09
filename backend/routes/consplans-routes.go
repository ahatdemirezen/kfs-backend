package routes

import (
	"kfs-backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupConsPlansRoutes(app *fiber.App) {
	consPlans := app.Group("/api/consplans")

	consPlans.Post("/", handlers.CreateConsPlan)
	consPlans.Get("/:consPlanId", handlers.GetConsPlan)
	consPlans.Get("/analysis/:analysisInfoId", handlers.GetConsPlansByAnalysisInfo)
	consPlans.Put("/:consPlanId", handlers.UpdateConsPlan)
	consPlans.Delete("/:consPlanId", handlers.DeleteConsPlan)

	// Birden fazla ConsPlan oluşturma
	consPlans.Post("/bulk", handlers.CreateMultipleConsPlans)
}
