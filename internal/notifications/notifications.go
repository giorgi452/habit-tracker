package notifications

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"habit-tracker/internal/habit"
)

func SendDiscordNotification(webhookURL string, h *habit.Habit) error {
	msg := fmt.Sprintf("**Habit Reminder**: %s\n**Estimated Time**: %s",
		h.Name, h.GetDurationReadable())

	payload := map[string]string{"content": msg}
	jsonBody, _ := json.Marshal(payload)

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
