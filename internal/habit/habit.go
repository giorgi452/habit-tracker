/* Package habit */
package habit

import (
	"fmt"
	"strings"
	"time"
)

type Habit struct {
	ID                int        `json:"id"`
	Name              string     `json:"name"`
	Freq              Frequency  `json:"freq"`
	Interval          int        `json:"interval"`
	Reminders         []Reminder `json:"reminders"`
	EstimatedDuration int        `json:"estimated_duration"`
	SpecificDate      *time.Time
	CreatedAt         time.Time  `json:"created_at"`
	LastCompleted     *time.Time `json:"last_completed"`
	IsArchived        bool
}

func AddHabit(name string, frequency Frequency, durationStr string, timeStrings []string) (*Habit, error) {
	h := &Habit{
		Name:      name,
		Freq:      frequency,
		Interval:  1,
		Reminders: make([]Reminder, 0, len(timeStrings)),
		CreatedAt: time.Now(),
	}

	if durationStr != "" {
		if err := h.SetDuration(durationStr); err != nil {
			return nil, err
		}
	}

	for _, ts := range timeStrings {
		cleaned := strings.TrimSpace(ts)
		if cleaned == "" {
			continue
		}
		if err := h.AddReminder(cleaned); err != nil {
			return nil, err
		}
	}
	return h, nil
}

func (h *Habit) String() string {
	var times []string
	for _, r := range h.Reminders {
		times = append(times, fmt.Sprintf("%02d:%02d", r.Hour, r.Minute))
	}

	status := "Pending"
	if h.LastCompleted != nil {
		status = fmt.Sprintf("Last done: %s", h.LastCompleted.Format("Jan 02"))
	}

	return fmt.Sprintf("[%d] %s | Every: %s | Duration: %ds | Times: %s | %s",
		h.ID, h.Name, h.Freq, h.EstimatedDuration, strings.Join(times, ","), status)
}
