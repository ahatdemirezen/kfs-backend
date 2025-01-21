package database

import (
	"fmt"
	"log"
	"os"

	"kfs-backend/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		os.Getenv("SUPABASE_HOST"),
		os.Getenv("SUPABASE_USER"),
		os.Getenv("SUPABASE_PASSWORD"),
		os.Getenv("SUPABASE_DBNAME"),
		os.Getenv("SUPABASE_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	log.Println("Veritabanına bağlantı başarılı")

	err = db.AutoMigrate(&models.User{}, &models.Profile{})
	if err != nil {
		log.Fatal("Veritabanı tabloları oluşturulamadı: ", err)
	}

	DB = db
}
