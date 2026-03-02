package habit

import "strings"

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
