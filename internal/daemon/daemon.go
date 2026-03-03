package daemon

import (
	"fmt"
	"time"

	"habit-tracker/internal/habit"
	"habit-tracker/internal/notifications"
)

func Start(webhookURL string, habits []*habit.Habit) {
	fmt.Printf("Daemon started at %s\n", time.Now().Format("15:04"))
	fmt.Printf("WebHook: %s\n", webhookURL)

	ticker := time.NewTicker(1 * time.Minute)

	for range ticker.C {
		now := time.Now()

		fmt.Printf("[%s] Scanning %d habits...\n", now.Format("15:04"), len(habits))

		for _, h := range habits {
			if h.IsDue() {
				fmt.Printf("Notification triggered for: %s\n", h.Name)

				err := notifications.SendDiscordNotification(webhookURL, h)
				if err != nil {
					fmt.Printf("Discord Error: %v\n", err)
				}
			}
		}
	}
}
