package menu

import (
	"bufio"
	"fmt"
	"strings"

	"habit-tracker/internal/habit"
)

func HandleDeleteHabit(scanner *bufio.Scanner, store *habit.Store) {
	h := ListAndSelectHabit(scanner, store)
	if h == nil {
		return
	}

	fmt.Printf("Are you sure you want to delete '%s'? (y/n): ", h.Name)
	scanner.Scan()
	if strings.ToLower(strings.TrimSpace(scanner.Text())) != "y" {
		fmt.Println("Delete cancelled.")
		return
	}

	if err := store.Delete(h.ID); err != nil {
		fmt.Printf("Error deleting habit: %v\n", err)
		return
	}
	fmt.Println("Habit deleted successfully.")
}
