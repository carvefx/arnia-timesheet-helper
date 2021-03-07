package main

import "time"

const development = "Development"
const publicHoliday = "Public Holiday"
const leave = "Leave"

type timesheetDay struct {
	date      time.Time
	projectID int
	status    dayStatus
}

type dayStatus struct {
	isWorked    bool
	description string
}

func newWorkDayStatus() *dayStatus {
	return &dayStatus{
		isWorked:    true,
		description: development,
	}

}

func newFreeDayStatus(kind string) *dayStatus {
	return &dayStatus{
		isWorked:    false,
		description: kind,
	}
}
