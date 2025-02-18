package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupMarketInfoRoutes(app *fiber.App) {
	marketInfoService := &services.GenericVentureService[models.MarketInfo]{} // HL

	// Route gruplarını olustur
	marketInfoGroup := app.Group("/api/market-infos")

	// Yeni MarketInfo olustur
	marketInfoGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.MarketInfoRequest
		return handlers.CreateVenture(c, marketInfoService, req)
	})

	// ID ile MarketInfo getir
	marketInfoGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, marketInfoService)
	})

	// Campaign ID'ye gore MarketInfo listele
	marketInfoGroup.Get("/list/:campaign_Id", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, marketInfoService, "campaign_id")
	})

	// MarketInfo guncelle
	marketInfoGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.MarketInfoRequest
		return handlers.UpdateVenture(c, marketInfoService, req)
	})

	// MarketInfo sil
	marketInfoGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, marketInfoService)
	})
}