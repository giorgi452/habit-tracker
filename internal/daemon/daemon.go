package daemon

import (
	"fmt"
	"sync"
	"time"

	"habit-tracker/internal/habit"
	"habit-tracker/internal/notifications"
)

func Start(webhookURL string, habits *[]*habit.Habit, mu *sync.RWMutex) {
	fmt.Printf("Daemon started at %s\n", time.Now().Format("15:04"))

	time.Sleep(time.Duration(60-time.Now().Second()) * time.Second)
	ticker := time.NewTicker(1 * time.Minute)

	for range ticker.C {
		mu.RLock()
		currentHabits := *habits

		for _, h := range currentHabits {
			if h.IsDue() {
				go func(target *habit.Habit) {
					err := notifications.SendDiscordNotification(webhookURL, target)
					if err != nil {
						fmt.Printf("Discord Error for %s: %v\n", target.Name, err)
					}
				}(h)
			}
		}
		mu.RUnlock()
	}
}
