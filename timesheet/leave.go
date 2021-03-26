package timesheet

import (
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

type Leave struct {
	days []int
}

// Parse the flag input from the main command
func (l *Leave) Parse(cfg Config, ldom int) {
	strLeaveDays := strings.Split(cfg.LeaveDays, ",")
	intLeaveDays := make([]int, 0)

	for _, s := range strLeaveDays {
		parsed, _ := strconv.Atoi(s)

		if parsed >= ldom || parsed <= 0 {
			continue
		}

		intLeaveDays = append(intLeaveDays, parsed)
	}

	l.days = intLeaveDays
}

func (l *Leave) HasDate(d int) bool {
	return funk.ContainsInt(l.days, d)
}

func (l *Leave) GetDates() []int {
	return l.days
}
