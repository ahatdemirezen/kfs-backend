// package routes

// import (
// 	"kfs-backend/handlers"

// 	"github.com/gofiber/fiber/v2"
// )

// // Kampanya rotalarını tanımla
// func SetupCampaignRoutes(app *fiber.App) {
// 	campaignGroup := app.Group("api/campaigns")

// 	campaignGroup.Post("/create-campaign", handlers.CreateCampaign)   // Yeni kampanya oluştur
// 	campaignGroup.Get("/get-campaign/:id", handlers.GetCampaignByID)  // ID ile kampanya getir
// 	campaignGroup.Get("/list-campaigns", handlers.GetAllCampaigns)    // Tüm kampanyaları getir
// 	campaignGroup.Put("/update-campaign/:id", handlers.UpdateCampaign) // Kampanyayı güncelle
// 	campaignGroup.Delete("/delete-campaign/:id", handlers.DeleteCampaign) // Kampanyayı sil
// }



package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

// Kampanya rotalarını tanımla
func SetupCampaignRoutes(app *fiber.App) {
	campaignGroup := app.Group("api/campaigns")

	// **JWT doğrulaması gerektirmeyen rotalar**
	campaignGroup.Get("/get-campaign/:id", handlers.GetCampaignByID)  // Belirli bir kampanyayı getir
	campaignGroup.Get("/list-campaigns", handlers.GetAllCampaigns)    // Tüm kampanyaları getir

	// **JWT doğrulaması gerektiren rotalar**
	campaignGroup.Use(middleware.AuthenticateMiddleware)

	campaignGroup.Post("/create-campaign", handlers.CreateCampaign)   // Yeni kampanya oluştur
	campaignGroup.Put("/update-campaign/:id", handlers.UpdateCampaign) // Kampanyayı güncelle
	campaignGroup.Delete("/delete-campaign/:id", handlers.DeleteCampaign) // Kampanyayı sil
}
