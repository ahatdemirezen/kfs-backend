package main

import (
	"log"

	"kfs-backend/config"
	"kfs-backend/database" // Veritabanı bağlantısı
	"kfs-backend/routes"   // Route'lar

	"github.com/gofiber/fiber/v2" // Fiber framework
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Fiber uygulamasını başlat
	app := fiber.New()

	// CORS middleware'ini ekle
	app.Use(cors.New())

	// Config yükle
	config.LoadConfig()

	// Veritabanı bağlantısını başlat
	database.ConnectDB()

	// Ana sayfa
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Fonbulucu API'sine Hoş Geldiniz")
	})

	// Tüm route'ları ayarla
	routes.SetupAuthRoutes(app)
	routes.SetupUserRoutes(app) // User routes'ları ekle

	// Debug için tüm route'ları yazdır
	for _, route := range app.GetRoutes() {
		log.Printf("Route: %s %s", route.Method, route.Path)
	}

	// Uygulamayı başlat
	log.Fatal(app.Listen(":3000"))
}
