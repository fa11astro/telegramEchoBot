package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func HandleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch {
	case update.Message.Text != "":
		// Эхо-ответ на текстовое сообщение
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		_, err := bot.Send(msg)
		if err != nil {
			log.Println("Ошибка при отправке сообщения:", err)
		}

	case update.Message.Sticker != nil:
		// Эхо-ответ на стикер
		msg := tgbotapi.NewSticker(update.Message.Chat.ID, tgbotapi.FileID(update.Message.Sticker.FileID))
		_, err := bot.Send(msg)
		if err != nil {
			log.Println("Ошибка при отправке стикера:", err)
		}

	case update.Message.Photo != nil:
		// Эхо-ответ на фото
		photo := update.Message.Photo[len(update.Message.Photo)-1] // Берем самое большое фото
		msg := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileID(photo.FileID))
		_, err := bot.Send(msg)
		if err != nil {
			log.Println("Ошибка при отправке фото:", err)
		}

	case update.Message.Voice != nil:
		// Эхо-ответ на голосовое сообщение
		msg := tgbotapi.NewVoice(update.Message.Chat.ID, tgbotapi.FileID(update.Message.Voice.FileID))
		_, err := bot.Send(msg)
		if err != nil {
			log.Println("Ошибка при отправке голосового сообщения:", err)
		}

	case update.Message.Video != nil:
		// Эхо-ответ на видео
		msg := tgbotapi.NewVideo(update.Message.Chat.ID, tgbotapi.FileID(update.Message.Video.FileID))
		_, err := bot.Send(msg)
		if err != nil {
			log.Println("Ошибка при отправке видео:", err)
		}

	case update.Message.VideoNote != nil:
		// Эхо-ответ на видеосообщение (кружочек)
		msg := tgbotapi.NewVideoNote(update.Message.Chat.ID, update.Message.VideoNote.Length, tgbotapi.FileID(update.Message.VideoNote.FileID))
		_, err := bot.Send(msg)
		if err != nil {
			log.Println("Ошибка при отправке видеосообщения:", err)
		}

	case update.Message.Document != nil:
		// Эхо-ответ на файл
		document := update.Message.Document

		// Проверяем, является ли файл музыкальным (по MIME-типу или расширению)
		if document.MimeType == "audio/mpeg" || document.MimeType == "audio/ogg" || document.FileName[len(document.FileName)-4:] == ".mp3" || document.FileName[len(document.FileName)-4:] == ".ogg" {
			// Отправляем аудиофайл
			msg := tgbotapi.NewAudio(update.Message.Chat.ID, tgbotapi.FileID(document.FileID))
			_, err := bot.Send(msg)
			if err != nil {
				log.Println("Ошибка при отправке аудиофайла:", err)
			}
		} else {
			// Отправляем файл как документ
			msg := tgbotapi.NewDocument(update.Message.Chat.ID, tgbotapi.FileID(document.FileID))
			_, err := bot.Send(msg)
			if err != nil {
				log.Println("Ошибка при отправке файла:", err)
			}
		}

	case update.Message.Audio != nil:
		// Эхо-ответ на аудио
		msg := tgbotapi.NewAudio(update.Message.Chat.ID, tgbotapi.FileID(update.Message.Audio.FileID))
		_, err := bot.Send(msg)
		if err != nil {
			log.Println("Ошибка при отправке аудио:", err)
		}

	case update.Message.Location != nil:
		// Эхо-ответ на местоположение
		location := update.Message.Location
		msg := tgbotapi.NewLocation(update.Message.Chat.ID, location.Latitude, location.Longitude)
		_, err := bot.Send(msg)
		if err != nil {
			log.Println("Ошибка при отправке местоположения:", err)
		}

	case update.Message.Poll != nil:
		// Эхо-ответ на опрос
		poll := update.Message.Poll

		// Извлекаем текстовые варианты ответов
		var options []string
		for _, option := range poll.Options {
			options = append(options, option.Text)
		}

		// Создаем новый опрос с теми же параметрами
		newPoll := tgbotapi.NewPoll(update.Message.Chat.ID, poll.Question, options...)
		newPoll.IsAnonymous = poll.IsAnonymous
		newPoll.Type = poll.Type
		newPoll.AllowsMultipleAnswers = poll.AllowsMultipleAnswers
		newPoll.CorrectOptionID = int64(poll.CorrectOptionID)

		// Отправляем опрос
		_, err := bot.Send(newPoll)
		if err != nil {
			log.Println("Ошибка при отправке опроса:", err)
		}

	case update.Message.Contact != nil:
		// Обработка номера телефона
		contact := update.Message.Contact

		// Отправляем контакт обратно в том же виде
		msg := tgbotapi.NewContact(update.Message.Chat.ID, contact.PhoneNumber, contact.FirstName)
		_, err := bot.Send(msg)
		if err != nil {
			log.Println("Ошибка при отправке контакта:", err)
		}

	default:
		// Ответ на неподдерживаемый тип сообщения
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я пока не умею обрабатывать этот тип сообщения 😅")
		_, err := bot.Send(msg)
		if err != nil {
			log.Println("Ошибка при отправке сообщения:", err)
		}
	}
}
