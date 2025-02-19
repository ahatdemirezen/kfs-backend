package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupThreatRoutes(app *fiber.App) {
	// Generic service'i Threat için başlat
	threatService := &services.GenericVentureService[models.Threat]{} // HL

	// Route gruplarını oluştur
	threatGroup := app.Group("/api/threats")

	// Yeni Threat oluştur
	threatGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.ThreatRequest
		return handlers.CreateVenture(c, threatService, req)
	})

	// ID ile Threat getir
	threatGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, threatService)
	})

	// AnalysisInfoId'ye göre Threat listele
	threatGroup.Get("/list/:analysis_info_id", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, threatService, "analysis_info_id")
	})

	// Threat güncelle
	threatGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.ThreatRequest
		return handlers.UpdateVenture(c, threatService, req)
	})

	// Threat sil
	threatGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, threatService)
	})
}