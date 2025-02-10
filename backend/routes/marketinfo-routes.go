package routes

import (
	"kfs-backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupMarketInfoRoutes(app *fiber.App) {
	marketInfos := app.Group("/api/marketinfos")

	marketInfos.Post("/create", handlers.CreateMarketInfo)
	marketInfos.Get("/:marketInfoId", handlers.GetMarketInfo)
	marketInfos.Get("/campaign/:campaignId", handlers.GetMarketInfosByCampaign)
	marketInfos.Put("/:marketInfoId", handlers.UpdateMarketInfo)
	marketInfos.Delete("/:marketInfoId", handlers.DeleteMarketInfo)
}
