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
	JwtSecret        string
	JwtSecretRefresh string
	NodeEnv          string
}

var AppConfig *Config

func LoadConfig(envPath string) {
	// .env dosyasını yükle
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file from path: %s", envPath)
	}

	// Config yapısını doldur
	AppConfig = &Config{
		SupabaseHost:     os.Getenv("SUPABASE_HOST"),
		SupabaseUser:     os.Getenv("SUPABASE_USER"),
		SupabasePassword: os.Getenv("SUPABASE_PASSWORD"),
		SupabaseDBName:   os.Getenv("SUPABASE_DBNAME"),
		SupabasePort:     os.Getenv("SUPABASE_PORT"),
		JwtSecret:        os.Getenv("JWT_SECRET"),
		JwtSecretRefresh: os.Getenv("JWT_SECRET_REFRESH"),
		NodeEnv:          os.Getenv("NODE_ENV"),
	}

	log.Println("Config dosyası başarıyla yüklendi")
}
