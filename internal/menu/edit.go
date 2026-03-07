package menu

import (
	"bufio"
	"fmt"
	"strings"

	"habit-tracker/internal/habit"
)

func HandleEditHabit(scanner *bufio.Scanner, store *habit.Store) {
	h := ListAndSelectHabit(scanner, store)
	if h == nil {
		return
	}

	fmt.Printf("Editing '%s' — press Enter to keep current value.\n", h.Name)

	fmt.Printf("Name [%s]: ", h.Name)
	scanner.Scan()
	newName := strings.TrimSpace(scanner.Text())

	fmt.Printf("Frequency [%s] (Daily/Weekly/Weekdays/Monthly): ", h.Freq)
	scanner.Scan()
	newFreqStr := strings.TrimSpace(scanner.Text())

	fmt.Printf("Duration [%ds] (e.g., 15m): ", h.EstimatedDuration)
	scanner.Scan()
	newDur := strings.TrimSpace(scanner.Text())

	err := store.Update(h.ID, func(target *habit.Habit) error {
		if newName != "" {
			target.Name = newName
		}
		if newFreqStr != "" {
			target.Freq = habit.ParseFrequency(newFreqStr)
		}
		if newDur != "" {
			if err := target.SetDuration(newDur); err != nil {
				return fmt.Errorf("Invalid duration: %w", err)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error updating habit: %v\n", err)
		return
	}
	fmt.Println("Habit updated successfully.")
}
