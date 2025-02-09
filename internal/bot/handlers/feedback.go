package handlers

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

func HandleFeedback(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	feedbackText := update.Message.CommandArguments()
	var msg tgbotapi.MessageConfig

	if feedbackText == "" {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста, напишите ваш отзыв после команды /feedback.")
	} else {
		// Отправляем отзыв пользователю
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Спасибо за ваш отзыв!")

		// Отправляем отзыв вам
		yourChatID := int64(1217039785) // chat_id
		feedbackMsg := tgbotapi.NewMessage(yourChatID, fmt.Sprintf("Новый отзыв от @%s:\n%s", update.Message.From.UserName, feedbackText))
		_, err := bot.Send(feedbackMsg)
		if err != nil {
			log.Println("Ошибка при отправке отзыва:", err)
		}

		// Сохраняем отзыв в файл
		file, err := os.OpenFile("feedback.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println("Ошибка при открытии файла:", err)
		} else {
			defer func(file *os.File) {
				err := file.Close()
				if err != nil {
					log.Println("Ошибка при закрытии файла:", err)
				}
			}(file)
			feedbackLine := fmt.Sprintf("Отзыв от @%s: %s\n", update.Message.From.UserName, feedbackText)
			if _, err := file.WriteString(feedbackLine); err != nil {
				log.Println("Ошибка при записи в файл:", err)
			}
		}
	}

	_, err := bot.Send(msg)
	if err != nil {
		log.Println("Ошибка при отправке сообщения:", err)
	}
}
