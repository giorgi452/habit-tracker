package habit

import (
	"fmt"
	"time"
)

func (h *Habit) SetDuration(input string) error {
	d, err := time.ParseDuration(input)
	if err != nil {
		return fmt.Errorf("invalid duration format (use 5m, 30s, etc): %v", err)
	}
	h.EstimatedDuration = int(d.Seconds())
	return nil
}

func (h *Habit) GetDurationReadable() string {
	d := time.Duration(h.EstimatedDuration) * time.Second
	return d.String()
}
