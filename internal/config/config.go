package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	DatabaseURL    string
	RedisURL       string
	JWTSecret      string
	TelegramToken  string
}

var App Config

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	App = Config{
		Port:          os.Getenv("PORT"),
		DatabaseURL:   os.Getenv("DATABASE_URL"),
		RedisURL:      os.Getenv("REDIS_URL"),
		JWTSecret:     os.Getenv("JWT_SECRET"),
		TelegramToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
	}
}