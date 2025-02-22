package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupFundingInfoRoutes(app *fiber.App) {
	// Generic service'i FundingInfo için başlat
	fundingInfoService := &services.GenericVentureService[models.FundingInfo]{} // HL

	// Route gruplarını oluştur
	fundingInfoGroup := app.Group("/api/funding-info")

	// Yeni FundingInfo oluştur
	fundingInfoGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.FundingInfoRequest
		return handlers.CreateVenture(c, fundingInfoService, req)
	})

	// ID ile FundingInfo getir
	fundingInfoGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, fundingInfoService)
	})

	// CampaignId'ye göre FundingInfo listele
	fundingInfoGroup.Get("/list/:campaign_Id", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, fundingInfoService, "campaign_id")
	})

	// FundingInfo güncelle
	fundingInfoGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.FundingInfoRequest
		return handlers.UpdateVenture(c, fundingInfoService, req)
	})

	// FundingInfo sil
	fundingInfoGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, fundingInfoService)
	})
}
