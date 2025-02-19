package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupConsRoutes(app *fiber.App) {
	// Generic service'i Cons için başlat
	consService := &services.GenericVentureService[models.Cons]{} // HL

	// Route gruplarını oluştur
	consGroup := app.Group("/api/cons")

	// Yeni Cons oluştur
	consGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.ConsRequest
		return handlers.CreateVenture(c, consService, req)
	})

	// ID ile Cons getir
	consGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, consService)
	})

	// AnalysisInfoId'ye göre Cons listele
	consGroup.Get("/list/:analysis_info_id", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, consService, "analysis_info_id")
	})

	// Cons güncelle
	consGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.ConsRequest
		return handlers.UpdateVenture(c, consService, req)
	})

	// Cons sil
	consGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, consService)
	})
}
