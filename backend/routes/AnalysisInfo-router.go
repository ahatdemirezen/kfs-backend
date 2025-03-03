package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupAnalysisInfoRoutes(app *fiber.App) {
	analysisInfoService := &services.GenericVentureService[models.AnalysisInfo]{} // HL

	// Route gruplarını olustur
	analysisInfoGroup := app.Group("/api/analysis-infos")

	// Yeni AnalysisInfo olustur
	analysisInfoGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.AnalysisInfoRequest
		return handlers.CreateVenture(c, analysisInfoService, req)
	})

	// ID ile AnalysisInfo getir
	analysisInfoGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, analysisInfoService)
	})

	// Campaign ID'ye gore AnalysisInfo listele
	analysisInfoGroup.Get("/list/:campaign_Id", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, analysisInfoService, "campaign_id")
	})

	// AnalysisInfo guncelle
	analysisInfoGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.AnalysisInfoRequest
		return handlers.UpdateVenture(c, analysisInfoService, req)
	})

	// AnalysisInfo sil
	analysisInfoGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, analysisInfoService)
	})

}
