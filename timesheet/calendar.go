package timesheet

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type Config struct {
	SelectedYear      int
	SelectedMonth     int
	SelectedProjectID int
	LeaveDays         string
	ArniaProjectID    int
}

type Calendar struct {
	leave         configDateParser
	weekend       configDateParser
	publicHoliday configDateParser
	logger        log.Logger
}

func NewCalendar(leave configDateParser, weekend configDateParser, publicHoliday configDateParser, l log.Logger) Calendar {
	return Calendar{
		leave:         leave,
		weekend:       weekend,
		publicHoliday: publicHoliday,
		logger:        l,
	}
}

func (c *Calendar) BuildTimesheet(cfg Config) []Day {
	currYear := cfg.SelectedYear
	currMonth := cfg.SelectedMonth

	ldom := c.getLastDayOfMonth(cfg.SelectedYear, cfg.SelectedMonth)
	c.leave.Parse(cfg, ldom)
	c.weekend.Parse(cfg, ldom)
	c.publicHoliday.Parse(cfg, ldom)

	level.Info(c.logger).Log("msg", "Processed dates as follows: ",
		"leaveDays", fmt.Sprintf("%d", c.leave.GetDates()),
		"publicHolidays", fmt.Sprintf("%d", c.publicHoliday.GetDates()),
		"weekendDays", fmt.Sprintf("%d", c.weekend.GetDates()))

	return c.dailyCheck(ldom, currYear, currMonth, cfg.ArniaProjectID, cfg.SelectedProjectID)
}

func (c *Calendar) getLastDayOfMonth(year int, month int) int {

	return time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.Local).Day()
}

func (c *Calendar) dailyCheck(ldom int, currYear int, currMonth int, arniaProjectID int, projectID int) []Day {
	timesheetDays := make([]Day, 0)
	for d := 1; d <= ldom; d++ {
		currDate := time.Date(currYear, time.Month(currMonth), d, 0, 0, 0, 0, time.Local)

		currTimeSheetDay := &Day{
			Date: currDate,
		}

		if c.leave.HasDate(d) {
			continue
		}

		if c.publicHoliday.HasDate(d) {
			currTimeSheetDay.Status = *newFreeDayStatus(publicHoliday)
			currTimeSheetDay.ProjectID = arniaProjectID
			timesheetDays = append(timesheetDays, *currTimeSheetDay)
			continue
		}

		if c.leave.HasDate(d) {
			currTimeSheetDay.Status = *newFreeDayStatus(leave)
			currTimeSheetDay.ProjectID = arniaProjectID
			timesheetDays = append(timesheetDays, *currTimeSheetDay)
			continue
		}

		currTimeSheetDay.Status = *newWorkDayStatus()
		currTimeSheetDay.ProjectID = projectID
		timesheetDays = append(timesheetDays, *currTimeSheetDay)
	}

	return timesheetDays
}
