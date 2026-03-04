package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	"habit-tracker/internal/daemon"
	"habit-tracker/internal/habit"
	"habit-tracker/internal/menu"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	webhook := os.Getenv("WEBHOOK_URL")

	var mu sync.RWMutex
	store := []*habit.Habit{}

	go daemon.Start(webhook, &store, &mu)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		op, err := menu.Print()
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch op {
		case 1:
			fmt.Print("Enter habit name: ")
			scanner.Scan()
			name := scanner.Text()

			fmt.Print("Enter frequency (Daily, Weekly, Weekdays): ")
			scanner.Scan()
			freq := habit.ParseFrequency(scanner.Text())

			fmt.Print("Duration (e.g. 15m): ")
			scanner.Scan()
			dur := scanner.Text()

			fmt.Print("Times (e.g. 09:00,17:00): ")
			scanner.Scan()
			tInput := strings.Split(scanner.Text(), ",")

			newHabit, err := habit.AddHabit(name, freq, dur, tInput)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			mu.Lock()
			store = append(store, newHabit)
			mu.Unlock()

			fmt.Println("✔ Habit added!")
		case 5:
			os.Exit(0)
			return
		}

	}
}
