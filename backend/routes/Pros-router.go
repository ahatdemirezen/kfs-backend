package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupProsRoutes(app *fiber.App) {
	// Generic service'i Pros için başlat
	prosService := &services.GenericVentureService[models.Pros]{} // HL

	// Route gruplarını oluştur
	prosGroup := app.Group("/api/pros")

	// Yeni Pros oluştur
	prosGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.ProsRequest
		return handlers.CreateVenture(c, prosService, req)
	})

	// ID ile Pros getir
	prosGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, prosService)
	})

	// AnalysisInfoId'ye göre Pros listele
	prosGroup.Get("/list/:analysis_info_id", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, prosService, "analysis_info_id")
	})

	// Pros güncelle
	prosGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.ProsRequest
		return handlers.UpdateVenture(c, prosService, req)
	})

	// Pros sil
	prosGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, prosService)
	})
}
