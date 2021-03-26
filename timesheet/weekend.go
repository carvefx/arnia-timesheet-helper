package timesheet

import (
	"time"

	"github.com/thoas/go-funk"
)

type Weekend struct {
	days []int
}

// ParseStr
func (w *Weekend) Parse(cfg Config, ldom int) {
	weekendDays := make([]int, 0)
	year := cfg.SelectedYear
	month := cfg.SelectedMonth

	for i := 1; i <= ldom; i++ {
		currDate := time.Date(year, time.Month(month), i, 0, 0, 0, 0, time.Local)

		if currDate.Weekday() == time.Saturday || currDate.Weekday() == time.Sunday {
			weekendDays = append(weekendDays, currDate.Day())
		}
	}

	w.days = weekendDays
}

func (w *Weekend) HasDate(d int) bool {
	return funk.ContainsInt(w.days, d)
}

func (l *Weekend) GetDates() []int {
	return l.days
}
