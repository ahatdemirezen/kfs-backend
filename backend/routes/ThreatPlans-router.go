package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupThreatPlanRoutes(app *fiber.App) {
	// Generic service'i ThreatPlan için başlat
	threatPlanService := &services.GenericVentureService[models.ThreatPlan]{} // HL

	// Route gruplarını oluştur
	threatPlanGroup := app.Group("/api/threat-plans")

	// Yeni ThreatPlan oluştur
	threatPlanGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.ThreatPlanRequest
		return handlers.CreateVenture(c, threatPlanService, req)
	})

	// ID ile ThreatPlan getir
	threatPlanGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, threatPlanService)
	})

	// AnalysisInfoId'ye göre ThreatPlan listele
	threatPlanGroup.Get("/list/:analysis_info_id", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, threatPlanService, "analysis_info_id")
	})

	// ThreatPlan güncelle
	threatPlanGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.ThreatPlanRequest
		return handlers.UpdateVenture(c, threatPlanService, req)
	})

	// ThreatPlan sil
	threatPlanGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, threatPlanService)
	})
}
