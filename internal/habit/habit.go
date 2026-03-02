/* Package habit */
package habit

import (
	"strings"
	"time"
)

type Frequency int

const (
	Once Frequency = iota
	Daily
	Weekly
	Monthly
	Weekdays
)

func (f Frequency) String() string {
	return [...]string{"Once", "Daily", "Weekly", "Monthly", "Weekdays"}[f]
}

func ParseFrequency(s string) Frequency {
	switch strings.ToLower(s) {
	case "daily":
		return Daily
	case "weekly":
		return Weekly
	case "monthly":
		return Monthly
	case "weekdays":
		return Weekdays
	default:
		return Once
	}
}

type Reminder struct {
	ID     int `json:"id"`
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

type Habit struct {
	ID            int        `json:"id"`
	Name          string     `json:"name"`
	Freq          Frequency  `json:"freq"`
	Interval      int        `json:"interval"`
	Reminders     []Reminder `json:"reminders"`
	SpecificDate  *time.Time
	CreatedAt     time.Time  `json:"created_at"`
	LastCompleted *time.Time `json:"last_completed"`
	IsArchived    bool
}

func (h *Habit) AddReminder(timeStr string) error {
	t, err := time.Parse("15:04", timeStr)
	if err != nil {
		return err
	}
	h.Reminders = append(h.Reminders, Reminder{Hour: t.Hour(), Minute: t.Minute()})
	return nil
}

func (h *Habit) SetDate(year, month, day int) {
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	h.SpecificDate = &date
}
