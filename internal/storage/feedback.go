package storage

import (
	"fmt"
	"os"
)

func SaveFeedback(username, feedback string) error {
	file, err := os.OpenFile("feedback.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	feedbackLine := fmt.Sprintf("Отзыв от @%s: %s\n", username, feedback)
	if _, err := file.WriteString(feedbackLine); err != nil {
		return err
	}

	return nil
}
