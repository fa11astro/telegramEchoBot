package bot

import (
	"fmt"
	"github.com/fa11astro/TelegramEchoBot/internal/bot/handlers"
	"github.com/fa11astro/TelegramEchoBot/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func Start(cfg *config.Config) error {
	// новый экземпляр бота
	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		return err
	}

	// вкл режим отладки
	bot.Debug = true

	log.Printf("Бот %s успешно запущен!", bot.Self.UserName)

	// вебхук
	_, err = bot.Request(tgbotapi.DeleteWebhookConfig{})
	if err != nil {
		return fmt.Errorf("ошибка при удалении вебхука: %v", err)
	}
	log.Println("Вебхук успешно удален, очередь сообщений очищена")

	// получение обновлений
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// входящие сообщения
	for update := range updates {
		if update.Message != nil {

			if update.Message.IsCommand() {
				handlers.HandleCommand(bot, update)
			} else {

				handlers.HandleMessage(bot, update)
			}
		}
	}

	return nil
}
