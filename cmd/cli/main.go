package main

import (
	"bufio"
	"fmt"
	"os"
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
			menu.HandleAddHabit(scanner, &store, &mu)
		case 5:
			return
		}
	}
}
