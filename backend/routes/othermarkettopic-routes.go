package routes

import (
	"kfs-backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupOtherMarketTopicRoutes(app *fiber.App) {
	topics := app.Group("/api/othermarkettopics")

	topics.Post("/", handlers.CreateOtherMarketTopic)
	topics.Get("/:topicId", handlers.GetOtherMarketTopic)
	topics.Get("/marketinfo/:marketInfoId", handlers.GetOtherMarketTopicsByMarketInfo)
	topics.Put("/:topicId", handlers.UpdateOtherMarketTopic)
	topics.Delete("/:topicId", handlers.DeleteOtherMarketTopic)
}
