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

	done := make(chan struct{})

	go func() {
		fmt.Printf("\nTracking '%s' (Target: %s)...\n", h.Name, h.GetDurationReadable())
		fmt.Println("Press 'd' + Enter to mark as done, or anything else to stop.")
		for {
			select {
			case t := <-ticker.C:
				elapsed := t.Sub(start).Round(time.Second)
				fmt.Printf("\rElapsed: %v", elapsed)
			case <-done:
				return
			}
		}
	}()

	scanner.Scan()
	close(done)
	input := strings.ToLower(strings.TrimSpace(scanner.Text()))

	elapsed := time.Since(start).Round(time.Second)
	fmt.Println()

	if input == "d" {
		now := time.Now()
		h.LastCompleted = &now
		fmt.Printf("Finished! Total time: %v\n", elapsed)
		if int(elapsed.Seconds()) < h.EstimatedDuration {
			fmt.Println("Great work! You finished faster than your estimate.")
		}
	} else {
		fmt.Println("Tracking stopped without completing.")
	}
}
