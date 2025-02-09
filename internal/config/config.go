package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	BotToken string
}

func Load() (*Config, error) {
	// Загружаем переменные из .env файла
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	// Получаем токен бота
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("Токен бота не указан в .env файле")
	}

	return &Config{
		BotToken: botToken,
	}, nil
}
