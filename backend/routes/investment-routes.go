package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupInvestmentRoutes(app *fiber.App) {
	investment := app.Group("/api/investments")

	// Kullanıcının kendi yatırımlarını oluşturması
	investment.Post("/:campaignId",
		middleware.AuthenticateMiddleware,
		middleware.IsInvestorMiddleware,
		handlers.CreateInvestment)

	// Kullanıcının kendi yatırımlarını görüntülemesi
	investment.Get("/my-investments",
		middleware.AuthenticateMiddleware,
		middleware.IsInvestorMiddleware,
		handlers.GetInvestments)

	// Tüm yatırımları görüntüleme (admin için)
	investment.Get("/all",
		middleware.AuthenticateMiddleware,
		middleware.AdminAuthMiddleware,
		handlers.GetAllInvestments)
}
