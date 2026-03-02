package habit

import "time"

func (h *Habit) SetDate(year, month, day int) {
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	h.SpecificDate = &date
}

func (h *Habit) IsDue() bool {
	now := time.Now()
	h24, m, _ := now.Clock()

	timeMatch := false
	for _, r := range h.Reminders {
		if r.Hour == h24 && r.Minute == m {
			timeMatch = true
			break
		}
	}
	if !timeMatch {
		return false
	}

	switch h.Freq {
	case Once:
		if h.IsArchived || h.SpecificDate == nil {
			return false
		}
		return h.SpecificDate.Format("2006-01-02") == now.Format("2006-01-02")
	case Weekdays:
		day := now.Weekday()
		if day == time.Saturday || day == time.Sunday {
			return false
		}
	}
	return true
}
