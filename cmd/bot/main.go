package main

import (
	"github.com/fa11astro/TelegramEchoBot/internal/bot"
	"github.com/fa11astro/TelegramEchoBot/internal/config"
	"log"
)

func main() {
	// Загружаем конфигурацию
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки конфигурации:", err)
	}

	// Запускаем бота
	if err := bot.Start(cfg); err != nil {
		log.Fatal("Ошибка при запуске бота:", err)
	}
}
