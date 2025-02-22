package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupFinancialDocumentsRoutes(app *fiber.App) {
	// Initialize generic service for FinancialDocuments
	financialDocumentsService := &services.GenericVentureService[models.FinancialDocuments]{}

	// Create route group
	financialDocumentsGroup := app.Group("/api/financial-documents")

	// Create new FinancialDocument
	financialDocumentsGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.FinancialDocumentsRequest
		return handlers.CreateVenture(c, financialDocumentsService, req)
	})

	// Retrieve FinancialDocument by ID
	financialDocumentsGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, financialDocumentsService)
	})

	// List FinancialDocuments by CampaignId
	financialDocumentsGroup.Get("/list/:campaign_Id", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, financialDocumentsService, "campaign_id")
	})

	// Update FinancialDocument
	financialDocumentsGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.FinancialDocumentsRequest
		return handlers.UpdateVenture(c, financialDocumentsService, req)
	})

	// Delete FinancialDocument
	financialDocumentsGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, financialDocumentsService)
	})
}
