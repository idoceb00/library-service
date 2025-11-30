package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	APIPort    string
}

func LoadConfig() *Config {
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file")
		} else {
			log.Println("Loaded .env file")
		}
	} else {
		log.Println("No .env file found; assuming database is already connected")
	}

	cfg := &Config{
		DBHost:     os.Getenv("PGHOST"),
		DBUser:     os.Getenv("PGUSER"),
		DBPassword: os.Getenv("PGPASSWORD"),
		DBName:     os.Getenv("PGDATABASE"),
		DBPort:     os.Getenv("PGPORT"),
		APIPort:    os.Getenv("API_PORT"),
	}

	if cfg.DBHost == "" || cfg.DBUser == "" || cfg.DBName == "" {
		log.Fatal("Missing required environment variables")
	}

	return cfg
}
