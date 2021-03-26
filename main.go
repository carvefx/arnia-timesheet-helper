package main

import (
	"github.com/carvefx/arnia-timesheet-helper/cli"
)

func main() {
	application := cli.NewApplication()
	application.Run()
	// timesheet.NewCa
	// timesheet := newTimesheet(*mFlag, *lFlag, *pFlag)
	// toCsv(timesheet.getEntries(), timesheet.currMonth)

}
