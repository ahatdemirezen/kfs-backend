package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupConsPlanRoutes(app *fiber.App) {
	// Generic service'i ConsPlan için başlat
	consPlanService := &services.GenericVentureService[models.ConsPlan]{} // HL

	// Route gruplarını oluştur
	consPlanGroup := app.Group("/api/cons-plans")

	// Yeni ConsPlan oluştur
	consPlanGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.ConsPlanRequest
		return handlers.CreateVenture(c, consPlanService, req)
	})

	// ID ile ConsPlan getir
	consPlanGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, consPlanService)
	})

	// AnalysisInfoId'ye göre ConsPlan listele
	consPlanGroup.Get("/list/:analysis_info_id", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, consPlanService, "analysis_info_id")
	})

	// ConsPlan güncelle
	consPlanGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.ConsPlanRequest
		return handlers.UpdateVenture(c, consPlanService, req)
	})

	// ConsPlan sil
	consPlanGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, consPlanService)
	})
}
