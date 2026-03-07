package menu

import (
	"bufio"
	"fmt"
	"strings"
	"time"

	"habit-tracker/internal/habit"
)

func StartHabit(scanner *bufio.Scanner, h *habit.Habit) {
	start := time.Now()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	done := make(chan bool)

	go func() {
		fmt.Printf("\nTracking '%s' (Target: %s)...\n", h.Name, h.GetDurationReadable())
		fmt.Println("Press 'd' + Enter to mark as completed.")

		for {
			select {
			case <-ticker.C:
				elapsed := time.Since(start).Round(time.Second)
				fmt.Printf("\rElapsed: %v", elapsed)
			case <-done:
				return
			}
		}
	}()

	scanner.Scan()
	input := strings.ToLower(strings.TrimSpace(scanner.Text()))
	done <- true

	if input == "d" {
		elapsed := time.Since(start)
		now := time.Now()
		h.LastCompleted = &now
		fmt.Printf("\nFinished! Total time: %v\n", elapsed.Round(time.Second))

		estimatedSeconds := h.EstimatedDuration
		if int(elapsed.Seconds()) < estimatedSeconds {
			fmt.Println("Great work! You finished faster than your estimate.")
		}
	} else {
		fmt.Println("\nTracking stopped without completing.")
	}
}
