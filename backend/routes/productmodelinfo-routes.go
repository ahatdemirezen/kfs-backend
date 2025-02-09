package routes

import (
	"kfs-backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupProductModelInfoRoutes(app *fiber.App) {
	infos := app.Group("/api/productmodelinfos")

	infos.Post("/create", handlers.CreateProductModelInfo)
	infos.Get("/:infoId", handlers.GetProductModelInfo)
	infos.Get("/campaign/:campaignId", handlers.GetProductModelInfosByCampaign)
	infos.Put("/:infoId", handlers.UpdateProductModelInfo)
	infos.Delete("/:infoId", handlers.DeleteProductModelInfo)
}
