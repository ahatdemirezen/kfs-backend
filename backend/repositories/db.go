package repositories

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
// try com
var DB *gorm.DB

func ConnectDatabase() {
	// .env dosyasını yükle
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Supabase DSN'i .env dosyasından al
	dsn := os.Getenv("SUPABASE_DSN")

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Veritabanına bağlanılamadı:", err)
	}

	log.Println("Veritabanına bağlantı başarılı")

	DB = database
}
//deneme