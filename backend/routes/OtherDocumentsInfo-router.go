package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupOtherDocumentsInfoRoutes(app *fiber.App) {
	// Initialize generic service for OtherDocumentsInfo
	otherDocumentsInfoService := &services.GenericVentureService[models.OtherDocumentsInfo]{}

	// Create route group
	otherDocumentsInfoGroup := app.Group("/api/other-documents-info")

	// Create new OtherDocumentsInfo
	otherDocumentsInfoGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.OtherDocumentsInfoRequest
		return handlers.CreateVenture(c, otherDocumentsInfoService, req)
	})

	// Retrieve OtherDocumentsInfo by ID
	otherDocumentsInfoGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, otherDocumentsInfoService)
	})

	// List OtherDocumentsInfo by CampaignId
	otherDocumentsInfoGroup.Get("/list/:campaign_Id", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, otherDocumentsInfoService, "campaign_id")
	})

	// Update OtherDocumentsInfo
	otherDocumentsInfoGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.OtherDocumentsInfoRequest
		return handlers.UpdateVenture(c, otherDocumentsInfoService, req)
	})

	// Delete OtherDocumentsInfo
	otherDocumentsInfoGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, otherDocumentsInfoService)
	})
}
