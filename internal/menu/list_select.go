package menu

import (
	"bufio"
	"fmt"
	"sync"

	"habit-tracker/internal/habit"
)

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
