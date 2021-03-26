package timesheet

import "time"

const development = "Development"
const publicHoliday = "Public Holiday"
const leave = "Leave"

type Day struct {
	Date      time.Time
	ProjectID int
	Status    dayStatus
}

type dayStatus struct {
	IsWorked    bool
	Description string
}

func newWorkDayStatus() *dayStatus {
	return &dayStatus{
		IsWorked:    true,
		Description: development,
	}

}

func newFreeDayStatus(kind string) *dayStatus {
	return &dayStatus{
		IsWorked:    false,
		Description: kind,
	}
}
