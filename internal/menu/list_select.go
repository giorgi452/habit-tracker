package menu

import (
	"bufio"
	"fmt"

	"habit-tracker/internal/habit"
)

func ListHabits(store *habit.Store) {
	habits := store.List()
	if len(habits) == 0 {
		fmt.Println("No habits found. Add one first!")
		return
	}

	fmt.Println("\n--- Your Habits ---")
	for _, h := range habits {
		fmt.Println(h)
	}
}

func ListAndSelectHabit(scanner *bufio.Scanner, store *habit.Store) *habit.Habit {
	habits := store.List()
	if len(habits) == 0 {
		fmt.Println("No habits found. Add one first!")
		return nil
	}

	fmt.Println("\n--- Your Habits ---")
	for _, h := range habits {
		fmt.Println(h)
	}

	fmt.Print("\nEnter Habit ID (or 0 to cancel): ")
	scanner.Scan()

	var id int
	fmt.Sscanf(scanner.Text(), "%d", &id)
	if id <= 0 {
		return nil
	}

	h, err := store.Get(id)
	if err != nil {
		fmt.Println("Invalid ID.")
		return nil
	}
	return h
}
