package cli

import (
	"os"

	"github.com/carvefx/arnia-timesheet-helper/csv"
	"github.com/carvefx/arnia-timesheet-helper/timesheet"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type Application struct {
	config   timesheet.Config
	calendar timesheet.Calendar
	logger   log.Logger
}

func NewApplication() *Application {
	logger := log.NewLogfmtLogger(os.Stdout)

	cfg, err := BuildConfigFromFlags()
	if err != nil {
		level.Error(logger).Log("err", err.Error())
		os.Exit(1)
	}

	leave := &timesheet.Leave{}
	weekend := &timesheet.Weekend{}
	api := timesheet.NewPublicHolidayAPI(logger)
	publicHoliday := timesheet.NewPublicHoliday(*api, logger)
	calendar := timesheet.NewCalendar(leave, weekend, publicHoliday, logger)

	return &Application{*cfg, calendar, logger}
}

func (app *Application) Run() {
	timesheet := app.calendar.BuildTimesheet(app.config)
	csvWriter := csv.NewWriter(timesheet, app.config.SelectedMonth, app.logger)

	csvWriter.Write()
}
