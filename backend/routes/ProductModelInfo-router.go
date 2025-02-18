package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupProductModelInfoRoutes(app *fiber.App) {
	// Generic service'i ProductModelInfo için baslat
	productModelInfoService := &services.GenericVentureService[models.ProductModelInfo]{} // HL

	// Route gruplarını olustur
	productModelInfoGroup := app.Group("/api/product-model-infos")

	// Yeni ProductModelInfo olustur
	productModelInfoGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.ProductModelInfoRequest
		return handlers.CreateVenture(c, productModelInfoService, req)
	})

	// ID ile ProductModelInfo getir
	productModelInfoGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, productModelInfoService)
	})

	// Campaign ID'ye gore ProductModelInfo listele
	productModelInfoGroup.Get("/list/:campaignId", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, productModelInfoService, "campaign_id")
	})

	// ProductModelInfo guncelle
	productModelInfoGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.ProductModelInfoRequest
		return handlers.UpdateVenture(c, productModelInfoService, req)
	})

	// ProductModelInfo sil
	productModelInfoGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, productModelInfoService)
	})
}