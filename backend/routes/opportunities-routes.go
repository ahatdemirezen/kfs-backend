package routes

import (
	"kfs-backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupOpportunitiesRoutes(app *fiber.App) {
	opportunities := app.Group("/api/opportunities")

	opportunities.Post("/", handlers.CreateOpportunity)
	opportunities.Get("/:opportunityId", handlers.GetOpportunity)
	opportunities.Get("/analysis/:analysisInfoId", handlers.GetOpportunitiesByAnalysisInfo)
	opportunities.Put("/:opportunityId", handlers.UpdateOpportunity)
	opportunities.Delete("/:opportunityId", handlers.DeleteOpportunity)

	// Birden fazla Opportunity oluşturma
	opportunities.Post("/bulk", handlers.CreateMultipleOpportunities)
}
