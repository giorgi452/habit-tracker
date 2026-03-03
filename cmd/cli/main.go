package main

import (
	"fmt"
	"log"
	"os"

	"habit-tracker/internal/daemon"
	"habit-tracker/internal/habit"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. Ensure it exists in the project root.")
	}

	h1, err := habit.AddHabit("Habit 1", habit.Daily, "1m", []string{"09:00", "21:00"})
	if err != nil {
		fmt.Println("Error creating habit 1:", err)
		return
	}

	h2, err := habit.AddHabit("Habit 2", habit.Weekdays, "1h", []string{"18:12"})
	if err != nil {
		fmt.Println("Error creating habit 2:", err)
		return
	}

	store := []*habit.Habit{h1, h2}

	daemon.Start(os.Getenv("WEBHOOK_URL"), store)
}
