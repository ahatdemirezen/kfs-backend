package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupPastCampaignInfoRoutes(app *fiber.App) {
	// Generic service'i PastCampaignInfo için başlat
	pastCampaignInfoService := &services.GenericVentureService[models.PastCampaignInfo]{}

	// Route gruplarını oluştur
	pastCampaignInfoGroup := app.Group("/api/past-campaign-infos")

	// Yeni PastCampaignInfo oluştur
	pastCampaignInfoGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.PastCampaignInfoRequest
		return handlers.CreateVenture(c, pastCampaignInfoService, req)
	})

	// ID ile PastCampaignInfo getir
	pastCampaignInfoGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, pastCampaignInfoService)
	})

	// Campaign ID'ye göre PastCampaignInfo listele
	pastCampaignInfoGroup.Get("/list/:campaignId", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, pastCampaignInfoService, "campaign_id")
	})

	// PastCampaignInfo güncelle
	pastCampaignInfoGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.PastCampaignInfoRequest
		return handlers.UpdateVenture(c, pastCampaignInfoService, req)
	})

	// PastCampaignInfo sil
	pastCampaignInfoGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, pastCampaignInfoService)
	})
}
