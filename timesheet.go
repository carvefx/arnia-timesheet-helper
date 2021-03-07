package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// timesheet describes a month in terms of workdays
type timesheet struct {
	currYear          int
	currMonth         int
	projectID         int
	leaveDays         []int
	publicHolidayDays []int
	weekendDays       []int
}

// getEntries returns the daily log with statuses
func (t *timesheet) getEntries() []timesheetDay {
	fmt.Println("Processed dates as follows:")
	fmt.Println("leaveDays", t.leaveDays, "publicHolidays", t.publicHolidayDays, "weekendDays", t.weekendDays)
	lastDay := getLastDayOfMonth(t.currYear, t.currMonth)

	calendar := make([]timesheetDay, 0)
	for d := 1; d <= lastDay; d++ {
		currDate := time.Date(t.currYear, time.Month(t.currMonth), d, 0, 0, 0, 0, time.Local)

		currTimeSheetDay := &timesheetDay{
			date: currDate,
		}

		if inDateRange(d, t.weekendDays) {
			continue
		}

		if inDateRange(d, t.publicHolidayDays) {
			currTimeSheetDay.status = *newFreeDayStatus(publicHoliday)
			currTimeSheetDay.projectID = arniaProjectID
			calendar = append(calendar, *currTimeSheetDay)
			continue
		}

		if inDateRange(d, t.leaveDays) {
			currTimeSheetDay.status = *newFreeDayStatus(leave)
			currTimeSheetDay.projectID = arniaProjectID
			calendar = append(calendar, *currTimeSheetDay)
			continue
		}

		currTimeSheetDay.status = *newWorkDayStatus()
		currTimeSheetDay.projectID = t.projectID
		calendar = append(calendar, *currTimeSheetDay)
	}

	return calendar
}

// is the current date in the array?
func inDateRange(num int, data []int) bool {
	if len(data) > 0 {
		for _, row := range data {
			if num == row {
				return true
			}
		}
	}
	return false
}

// NewTimesheet instantiates a new Timesheet struct
func newTimesheet(month int, leaveDays string, projectID int) *timesheet {
	year := time.Now().Year()

	return &timesheet{
		currYear:          year,
		currMonth:         month,
		projectID:         projectID,
		leaveDays:         parseLeaveDays(year, month, leaveDays),
		publicHolidayDays: getHolidaysByMonth(year, month),
		weekendDays:       parseWeekendsByMonth(year, month),
	}
}

// Parse the flag input from the main command
func parseLeaveDays(year int, month int, leaveDays string) []int {
	strLeaveDays := strings.Split(leaveDays, ",")
	intLeaveDays := make([]int, 0)

	for _, s := range strLeaveDays {
		parsed, _ := strconv.Atoi(s)

		if parsed >= getLastDayOfMonth(year, month) || parsed <= 0 {
			continue
		}

		intLeaveDays = append(intLeaveDays, parsed)
	}

	return intLeaveDays
}
