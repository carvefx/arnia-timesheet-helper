package timesheet

import (
	"os"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/thoas/go-funk"
)

type PublicHoliday struct {
	days   []int
	api    PublicHolidayAPI
	logger log.Logger
}

func NewPublicHoliday(api PublicHolidayAPI, logger log.Logger) *PublicHoliday {
	return &PublicHoliday{
		api:    api,
		logger: logger,
	}
}

// Parse processes the data from the API into a slice of day-dates of public holidays
func (ph *PublicHoliday) Parse(cfg Config, ldom int) {
	var filteredHolidays []int = make([]int, 0)
	year := cfg.SelectedYear
	month := cfg.SelectedMonth
	ignore := cfg.IgnorePublicHoliday

	if ignore {
	    return
	}

	h := ph.api.Fetch(year)

	for _, day := range h {
		currDate, err := time.Parse("2006-01-02", day.Date)

		if err != nil {
			level.Error(ph.logger).Log("err", err.Error())
			os.Exit(1)
		}

		if int(currDate.Month()) == month {
			filteredHolidays = append(filteredHolidays, currDate.Day())
		}
	}

	ph.days = filteredHolidays
}

func (ph *PublicHoliday) HasDate(d int) bool {
	return funk.ContainsInt(ph.days, d)
}

func (l *PublicHoliday) GetDates() []int {
	return l.days
}
