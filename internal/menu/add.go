package menu

import (
	"bufio"
	"fmt"
	"strings"

	"habit-tracker/internal/habit"
)

func HandleAddHabit(scanner *bufio.Scanner, store *habit.Store) {
	fmt.Print("Enter habit name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("Enter frequency (Daily, Weekly, Weekdays, Monthly): ")
	scanner.Scan()
	freq := habit.ParseFrequency(scanner.Text())

	fmt.Print("Duration (e.g., 15m, 30s): ")
	scanner.Scan()
	dur := scanner.Text()

	fmt.Print("Times (comma-separated, e.g., 09:00,17:00): ")
	scanner.Scan()
	times := strings.Split(scanner.Text(), ",")

	h, err := store.Add(name, freq, dur, times)
	if err != nil {
		fmt.Printf("Error creating habit: %v\n", err)
		return
	}

	fmt.Printf("Habit '%s' added with ID %d.\n", h.Name, h.ID)
}
