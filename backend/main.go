package main

import (
	"log"

	"kfs-backend/config"
	"kfs-backend/database" // Veritabanı bağlantısı
	"kfs-backend/middleware"
	"kfs-backend/routes"

	"github.com/gofiber/fiber/v2" // Fiber framework
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Fiber uygulamasını başlat
	app := fiber.New(fiber.Config{
		BodyLimit:    50 * 1024 * 1024, // 50MB'a yükseltiliyor
		ErrorHandler: middleware.ErrorHandler,
	})

	// CORS middleware'ini ekle
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3001", // Frontend URL'ini ekle
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true, // withCredentials için gerekli!
	}))

	// Config yükle
	config.LoadConfig()

	// Veritabanı bağlantısını başlat
	database.ConnectDB()

	// Ana sayfa
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Fonbulucu API'sine Hoş Geldiniz")
	})

	// Tüm route'ları ayarla
	routes.SetupRegisterRoutes(app)
	routes.SetupUserRoutes(app)            // User routes'ları ekle
	routes.SetupProfileRoutes(app)         // Profil rotaları
	routes.SetupAuthRoutes(app)            // auth route'larını kaydet
	routes.SetupVerificationRoutes(app)    // verification route'larını kaydet
	routes.SetupInvestmentRoutes(app)      // investment route'larını kaydet
	routes.SetupAdminRoutes(app)           // superAdmin route'larını kaydet
	routes.SetupRoleApplicationRoutes(app) // roleApp route'larını kaydet
	routes.SetupRoleappRoutes(app)         // roleApp route'larını kaydet
	routes.SetupAuthRoutes(app)
	routes.SetupUserRoutes(app) // User routes'ları ekle
	routes.SetupProfileRoutes(app)

	// Debug için tüm route'ları yazdır
	for _, route := range app.GetRoutes() {
		log.Printf("Route: %s %s", route.Method, route.Path)
	}

	// Campaign route'larını tanımla
	routes.SetupCampaignRoutes(app)

	// VentureCategory route'larını tanımla
	routes.SetupVentureCategoryRoutes(app)

	// VentureLocation route'larını tanımla
	routes.SetupVentureLocationRoutes(app)

	// VentureSector route'larını tanımla
	routes.SetupVentureSectorRoutes(app)

	// ParticipantEmail route'larını tanımla
	routes.SetupParticipantEmailRoutes(app)

	// VentureBusiness route'larını tanımla
	routes.SetupVentureBusinessModalRoutes(app)

	// VentureBusiness route'larını tanımla
	routes.SetupPastCampaignInfoRoutes(app)

	// Patent route'larını tanımla
	routes.SetupPatentRoutes(app)

	// Achivements route'larını tanımla
	routes.SetupAchivementsRoutes(app)

	// Permissions route'larını tanımla
	routes.SetupPermissionsRoutes(app)

	// TeamMembers route'larını tanımla
	routes.SetupTeamMembersRoutes(app)

	// ProductModelInfo route'larını tanımla
	routes.SetupProductModelInfoRoutes(app)

	// OtherProductTopics route'larını tanımla
	routes.SetupOtherProductTopicsRoutes(app)

	// MarketInfo route'larını tanımla
	routes.SetupMarketInfoRoutes(app)

	// OtherMarketTopics route'larını tanımla
	routes.SetupOtherMArketTopicsRoutes(app)

	// AnalysisInfo route'larını tanımla
	routes.SetupAnalysisInfoRoutes(app)

	// Pros route'larını tanımla
	routes.SetupProsRoutes(app)

	// Cons route'larını tanımla
	routes.SetupConsRoutes(app)

	// Opportunities route'larını tanımla
	routes.SetupOpportunityRoutes(app)

	// Threats route'larını tanımla
	routes.SetupThreatRoutes(app)

	// ConsPlans route'larını tanımla
	routes.SetupConsPlanRoutes(app)

	// ThreatPlans route'larını tanımla
	routes.SetupThreatPlanRoutes(app)

	// RisksInfo route'larını tanımla
	routes.SetupRisksInfoRoutes(app)

	// FundingInfo route'larını tanımla
	routes.SetupFundingInfoRoutes(app)

	// Usages route'larını tanımla
	routes.SetupUsageRoutes(app)

	// ExtraFinancingResources route'larını tanımla
	routes.SetupExtraFinancingResourcesRoutes(app)

	// EnterpriseInfo route'larını tanımla
	routes.SetupEnterpriseInfoRoutes(app)

	// AfterFundingFounderPartners route'larını tanımla
	routes.SetupAfterFundingFounderPartnersRoutes(app)

	// Videos route'larını tanımla
	routes.SetupVideosRoutes(app)

	// VisualInfo route'larını tanımla
	routes.SetupVisualInfoRoutes(app)

	// OtherPhotos route'larını tanımla
	routes.SetupOtherPhotosRoutes(app)

	// OtherDocumentsInfo route'larını tanımla
	routes.SetupOtherDocumentsInfoRoutes(app)

	// ProfitForecast route'larını tanımla
	routes.SetupProfitForecastRoutes(app)

	// InvestmentBudgets route'larını tanımla
	routes.SetupInvestmentBudgetsRoutes(app)

	// IncomeItems route'larını tanımla
	routes.SetupIncomeItemsRoutes(app)

	// FinancialDocuments route'larını tanımla
	routes.SetupFinancialDocumentsRoutes(app)

	// Explanations route'larını tanımla
	routes.SetupExplanationsRoutes(app)

	// SaleGoals route'larını tanımla
	routes.SetupSaleGoalsRoutes(app)

	// FinancialCategories route'larını tanımla
	routes.SetupFinancialCategoryRoutes(app)

	// FinancialExpenses route'larını tanımla
	routes.SetupFinancialExpenseRoutes(app)

	// FinancialSubCategories route'larını tanımla
	routes.SetupFinancialSubCategoryRoutes(app)

	// Venture route'larını tanımla
	routes.SetupGeneralRoutes(app)

	// Uygulamayı başlat
	log.Fatal(app.Listen(":3000"))
}
