package menu

import (
	"bufio"
	"fmt"
	"strings"
	"sync"
	"time"

	"habit-tracker/internal/habit"
)

func HandleAddHabit(scanner *bufio.Scanner, store *[]*habit.Habit, mu *sync.RWMutex) {
	fmt.Print("Enter habit name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("Enter frequency (Daily, Weekly, Weekdays, Monthly): ")
	scanner.Scan()
	freqStr := scanner.Text()
	freq := habit.ParseFrequency(freqStr)

	fmt.Print("Duration (e.g., 15m, 30s): ")
	scanner.Scan()
	dur := scanner.Text()

	fmt.Print("Times (comma separated, e.g., 09:00, 17:00): ")
	scanner.Scan()
	tInput := strings.Split(scanner.Text(), ",")

	newHabit, err := habit.AddHabit(name, freq, dur, tInput)
	if err != nil {
		fmt.Printf("Error creating habit: %v\n", err)
		return
	}

	mu.Lock()
	newHabit.ID = len(*store) + 1
	*store = append(*store, newHabit)
	mu.Unlock()

	fmt.Println("Habit added successfully!")
}

func ListAndSelectHabit(scanner *bufio.Scanner, store *[]*habit.Habit, mu *sync.RWMutex) *habit.Habit {
	mu.RLock()
	defer mu.RUnlock()

	if len(*store) == 0 {
		fmt.Println("No habits found. Add one first!")
		return nil
	}

	fmt.Println("\n--- Your Habits ---")
	for _, h := range *store {
		fmt.Println(h.String())
	}

	fmt.Print("\nEnter Habit ID to select (or 0 to cancel): ")
	scanner.Scan()
	var selectedID int
	fmt.Sscanf(scanner.Text(), "%d", &selectedID)

	if selectedID <= 0 {
		return nil
	}

	for _, h := range *store {
		if h.ID == selectedID {
			return h
		}
	}

	fmt.Println("Invalid ID.")
	return nil
}

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
