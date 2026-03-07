package menu

import (
	"bufio"
	"fmt"
	"strings"
	"sync"

	"habit-tracker/internal/habit"
)

func HandleDeleteHabit(scanner *bufio.Scanner, store *[]*habit.Habit, mu *sync.RWMutex) {
	h := ListAndSelectHabit(scanner, store, mu)
	if h == nil {
		return
	}

	fmt.Printf("Are you sure you want to delete '%s'? (y/n): ", h.Name)
	scanner.Scan()
	confirm := strings.ToLower(strings.TrimSpace(scanner.Text()))

	if confirm == "y" {
		mu.Lock()
		defer mu.Unlock()

		for i, item := range *store {
			if item.ID == h.ID {
				*store = append((*store)[:i], (*store)[i+1:]...)
				fmt.Println("Habit deleted successfully.")
				return
			}
		}
	} else {
		fmt.Println("Delete cancelled.")
	}
}
