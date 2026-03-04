package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"habit-tracker/internal/daemon"
	"habit-tracker/internal/habit"
	"habit-tracker/internal/menu"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. Ensure it exists in the project root.")
	}

	var mu sync.Mutex
	store := []*habit.Habit{}

	go daemon.Start(os.Getenv("WEBHOOK_URL"), store)

	for {
		op, _ := menu.Print()

		switch op {
		case 1:
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Print("Enter habit name: ")
			scanner.Scan()
			name := scanner.Text()

			fmt.Print("Enter repeat frequency (Daily, Weekly): ")
			scanner.Scan()
			repeat := scanner.Text()

			fmt.Print("Enter duration (e.g., 30m): ")
			scanner.Scan()
			duration := scanner.Text()

			fmt.Print("Enter times (comma separated, e.g., 12:00,18:00): ")
			scanner.Scan()
			timesInput := scanner.Text()
			rawTimes := strings.Split(timesInput, ",")
			var times []string
			for _, t := range rawTimes {
				trimmed := strings.TrimSpace(t)
				if trimmed != "" {
					times = append(times, trimmed)
				}
			}

			newHabit, err := habit.AddHabit(name, habit.ParseFrequency(repeat), duration, times)
			if err != nil {
				fmt.Println("Error adding habit:", err)
				return
			}
			mu.Lock()
			store = append(store, newHabit)
			mu.Unlock()

			fmt.Println("Habit added successfully!")
		case 5:
			os.Exit(0)
			return
		}

	}
}
