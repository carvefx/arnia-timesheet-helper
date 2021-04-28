package cli

import (
	"errors"
	"flag"
	"time"

	"github.com/carvefx/arnia-timesheet-helper/timesheet"
)

const unspecifiedProjectID int = 000
const arniaProjectID int = 40
const ignorePublicHolidays bool = false

func BuildConfigFromFlags() (*timesheet.Config, error) {

	defaultYear := int(time.Now().Year())
	defaultMonth := int(time.Now().Month())
	lFlag := flag.String("l", "", "Comma-separated list of days where you took a Leave. e.g. 02,05,22")
	pFlag := flag.Int("p", unspecifiedProjectID, "Your project ID")
	mFlag := flag.Int("m", defaultMonth, "The month you want to generate the export against. Will default to the current month e.g. 06")
	iFlag := flag.Bool("i", ignorePublicHolidays, "Ignore public holidays <boolean field>")

	flag.Parse()

	if *pFlag == unspecifiedProjectID {
		return nil, errors.New("missing project ID")
	}

	if *pFlag == arniaProjectID {
		return nil, errors.New("invalid project ID. reserved for Arnia")
	}

	return &timesheet.Config{
		SelectedYear:        defaultYear,
		SelectedMonth:       *mFlag,
		SelectedProjectID:   *pFlag,
		LeaveDays:           *lFlag,
		IgnorePublicHoliday: *iFlag,
		ArniaProjectID:      arniaProjectID,
	}, nil
}
