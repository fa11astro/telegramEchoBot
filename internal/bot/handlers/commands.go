package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func HandleCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Я эхо-бот. Отправь мне сообщение, и я повторю его. Используй /help для списка команд.")
		_, err := bot.Send(msg)
		if err != nil {
			log.Println("Ошибка при отправке сообщения:", err)
		}

	case "help":
		helpText := `Доступные команды:
/start - Начать диалог с ботом
/help - Показать список команд
/about - Информация о боте
/feedback <текст> - Оставить отзыв`
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, helpText)
		_, err := bot.Send(msg)
		if err != nil {
			log.Println("Ошибка при отправке сообщения:", err)
		}

	case "about":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я эхо-бот бла бла бла... повторяю сообщения и бла бла бла")
		_, err := bot.Send(msg)
		if err != nil {
			log.Println("Ошибка при отправке сообщения:", err)
		}

	case "feedback":
		HandleFeedback(bot, update) // Выносим обработку фидбэка в отдельную функцию

	default:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда. Используй /help для списка команд.")
		_, err := bot.Send(msg)
		if err != nil {
			log.Println("Ошибка при отправке сообщения:", err)
		}
	}
}
