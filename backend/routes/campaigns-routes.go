package routes

import (
	"kfs-backend/handlers"

	"github.com/gofiber/fiber/v2"
)

// Kampanya rotalarını tanımla
func SetupCampaignRoutes(app *fiber.App) {
	campaignGroup := app.Group("api/campaigns")

	campaignGroup.Post("/create-campaign", handlers.CreateCampaign)   // Yeni kampanya oluştur
	campaignGroup.Get("/get-campaign/:id", handlers.GetCampaignByID)  // ID ile kampanya getir
	campaignGroup.Get("/list-campaigns", handlers.GetAllCampaigns)    // Tüm kampanyaları getir
	campaignGroup.Put("/update-campaign/:id", handlers.UpdateCampaign) // Kampanyayı güncelle
	campaignGroup.Delete("/delete-campaign/:id", handlers.DeleteCampaign) // Kampanyayı sil
}