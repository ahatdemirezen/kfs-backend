package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SupabaseHost     string
	SupabaseUser     string
	SupabasePassword string
	SupabaseDBName   string
	SupabasePort     string
}

var AppConfig *Config

func LoadConfig() {
	// .env dosyasını yükle
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Config yapısını doldur
	AppConfig = &Config{
		SupabaseHost:     os.Getenv("SUPABASE_HOST"),
		SupabaseUser:     os.Getenv("SUPABASE_USER"),
		SupabasePassword: os.Getenv("SUPABASE_PASSWORD"),
		SupabaseDBName:   os.Getenv("SUPABASE_DBNAME"),
		SupabasePort:     os.Getenv("SUPABASE_PORT"),
	}

	log.Println("Config dosyası başarıyla yüklendi")
}
