package menu

import (
	"bufio"
	"fmt"
	"strings"
	"sync"

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
		fmt.Printf("⚠️ Error creating habit: %v\n", err)
		return
	}

	mu.Lock()
	*store = append(*store, newHabit)
	mu.Unlock()

	fmt.Println("✔ Habit added successfully!")
}
