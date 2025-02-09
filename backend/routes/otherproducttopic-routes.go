package routes

import (
	"kfs-backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupOtherProductTopicRoutes(app *fiber.App) {
	topics := app.Group("/api/otherproducttopics")

	topics.Post("/", handlers.CreateOtherProductTopic)
	topics.Get("/:topicId", handlers.GetOtherProductTopic)
	topics.Get("/productmodelinfo/:productModelInfoId", handlers.GetOtherProductTopicsByProductModelInfo)
	topics.Put("/:topicId", handlers.UpdateOtherProductTopic)
	topics.Delete("/:topicId", handlers.DeleteOtherProductTopic)
}
