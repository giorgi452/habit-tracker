package habit

import "time"

type Reminder struct {
	ID     int `json:"id"`
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

func (h *Habit) AddReminder(timeStr string) error {
	t, err := time.Parse("15:04", timeStr)
	if err != nil {
		return err
	}
	h.Reminders = append(h.Reminders, Reminder{Hour: t.Hour(), Minute: t.Minute()})
	return nil
}
