package main

import (
	"bufio"
	"fmt"
	"os"

	"habit-tracker/internal/daemon"
	"habit-tracker/internal/habit"
	"habit-tracker/internal/menu"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	webhook := os.Getenv("WEBHOOK_URL")

	store := habit.NewStore()
	go daemon.Start(webhook, store)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		op, err := menu.Print()
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch op {
		case 1:
			menu.HandleAddHabit(scanner, store)
		case 2:
			if h := menu.ListAndSelectHabit(scanner, store); h != nil {
				menu.StartHabit(scanner, h)
			}
		case 3:
			menu.HandleEditHabit(scanner, store)
		case 4:
			menu.HandleDeleteHabit(scanner, store)
		case 5:
			menu.ListHabits(store)
		case 6:
			return
		}
	}
}
