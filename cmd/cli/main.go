package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"habit-tracker/internal/daemon"
	"habit-tracker/internal/habit"
	"habit-tracker/internal/menu"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. Ensure it exists in the project root.")
	}

	store := []*habit.Habit{}

	go daemon.Start(os.Getenv("WEBHOOK_URL"), store)

	for {
		menu.Print()
	}
}
