package main

import (
	"time"
)

// return a slice of weekend days for a given month
func parseWeekendsByMonth(year int, month int) []int {
	weekendDays := make([]int, 0)

	ld := getLastDayOfMonth(year, month)

	for i := 1; i <= ld; i++ {
		currDate := time.Date(year, time.Month(month), i, 0, 0, 0, 0, time.Local)

		if currDate.Weekday() == time.Saturday || currDate.Weekday() == time.Sunday {
			weekendDays = append(weekendDays, currDate.Day())
		}
	}

	return weekendDays
}

func getLastDayOfMonth(year int, month int) int {

	return time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.Local).Day()
}
