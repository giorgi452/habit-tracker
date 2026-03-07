package notifications

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"habit-tracker/internal/habit"
)

var httpClient = &http.Client{
	Timeout: 5 * time.Second,
}

func SendDiscordNotification(webhookURL string, h *habit.Habit) error {
	msg := fmt.Sprintf("**Habit Reminder**: %s\n**Estimated Time**: %s",
		h.Name, h.GetDurationReadable())

	payload := map[string]string{"content": msg}
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("Failed to marshal payload: %w", err)
	}

	resp, err := httpClient.Post(webhookURL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("Discord returned status %d", resp.StatusCode)
	}
	return nil
}
