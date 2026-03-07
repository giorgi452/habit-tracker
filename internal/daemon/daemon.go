package daemon

import (
	"fmt"
	"sync"
	"time"

	"habit-tracker/internal/habit"
	"habit-tracker/internal/notifications"
)

const maxWorkers = 5

func Start(webhookURL string, store *habit.Store) {
	fmt.Printf("Daemon started at %s\n", time.Now().Format("15:04"))

	time.Sleep(time.Duration(60-time.Now().Second()) * time.Second)
	ticker := time.NewTicker(time.Minute)

	for range ticker.C {
		dispatch(webhookURL, store)
	}
}

func dispatch(webhookURL string, store *habit.Store) {
	due := collectDue(store)
	if len(due) == 0 {
		return
	}

	jobs := make(chan *habit.Habit, len(due))
	for _, h := range due {
		jobs <- h
	}
	close(jobs)

	workers := min(maxWorkers, len(due))
	var wg sync.WaitGroup
	wg.Add(workers)

	for range workers {
		go func() {
			defer wg.Done()
			for h := range jobs {
				if err := notifications.SendDiscordNotification(webhookURL, h); err != nil {
					fmt.Printf("Discord error for '%s': %v\n", h.Name, err)
				}
			}
		}()
	}

	wg.Wait()
}

func collectDue(store *habit.Store) []*habit.Habit {
	var due []*habit.Habit
	store.Range(func(h *habit.Habit) {
		if h.IsDue() {
			due = append(due, h)
		}
	})
	return due
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
