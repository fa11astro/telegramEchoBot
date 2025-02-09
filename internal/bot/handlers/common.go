package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	if update.Message.IsCommand() {
		HandleCommand(bot, update) // Обработка команд
	} else {
		HandleMessage(bot, update) // Обработка обычных сообщений
	}
}
