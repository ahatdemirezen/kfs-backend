package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupOpportunityRoutes(app *fiber.App) {
	// Generic service'i Opportunity için başlat
	opportunityService := &services.GenericVentureService[models.Opportunity]{} // HL

	// Route gruplarını oluştur
	opportunityGroup := app.Group("/api/opportunities")

	// Yeni Opportunity oluştur
	opportunityGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.OpportunityRequest
		return handlers.CreateVenture(c, opportunityService, req)
	})

	// ID ile Opportunity getir
	opportunityGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, opportunityService)
	})

	// AnalysisInfoId'ye göre Opportunity listele
	opportunityGroup.Get("/list/:analysis_info_id", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, opportunityService, "analysis_info_id")
	})

	// Opportunity güncelle
	opportunityGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.OpportunityRequest
		return handlers.UpdateVenture(c, opportunityService, req)
	})

	// Opportunity sil
	opportunityGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, opportunityService)
	})
}
