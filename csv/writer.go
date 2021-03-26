package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/carvefx/arnia-timesheet-helper/timesheet"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type Writer struct {
	fileHeader []string
	timesheet  []timesheet.Day
	month      int
	logger     log.Logger
}

var fileHeader = []string{"date", "account_id/id", "journal_id/id", "name", "unit_amount"}

func NewWriter(timesheet []timesheet.Day, month int, logger log.Logger) *Writer {
	return &Writer{
		fileHeader: fileHeader,
		timesheet:  timesheet,
		month:      month,
		logger:     logger,
	}
}

func (c *Writer) Write() {
	currDate := time.Now()
	fileSuffix := fmt.Sprintf("%d-%d", currDate.Year(), c.month)
	file, err := os.Create("storage/timesheet-" + fileSuffix + ".csv")
	if err != nil {
		level.Error(c.logger).Log("err", "cannot create file", err)
		os.Exit(1)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(fileHeader)
	if err != nil {
		level.Error(c.logger).Log("err", "cannot write to file", err)
		os.Exit(1)
	}

	for _, entry := range c.timesheet {
		err := writer.Write(c.parseEntryIntoRow(entry))
		if err != nil {
			level.Error(c.logger).Log("err", "cannot write to file", err)
			os.Exit(1)
		}
	}
}

func (c *Writer) parseEntryIntoRow(entry timesheet.Day) []string {
	row := make([]string, 5)

	row[0] = entry.Date.Format("2006-01-02")
	row[1] = "__export__.account_analytic_account_" + fmt.Sprint(entry.ProjectID)
	row[2] = "hr_timesheet.analytic_journal"
	row[3] = entry.Status.Description
	row[4] = "8"

	return row
}
