package main

import (
	"flag"
	"log"
	"time"
)

const arniaProjectID int = 40
const unspecifiedProjectID int = 000

func main() {
	defaultMonth := int(time.Now().Month())
	var lFlag = flag.String("l", "", "Comma-separated list of days where you took a Leave. e.g. 02,05,22")
	var pFlag = flag.Int("p", unspecifiedProjectID, "Your project id")
	var mFlag = flag.Int("m", defaultMonth, "The month you want to generate the export against. Will default to the current month e.g. 06")

	flag.Parse()

	if *pFlag == unspecifiedProjectID {
		log.Fatal("You need to specify a valid project id. You can find it in Odoo under analyticID")
	}

	timesheet := newTimesheet(*mFlag, *lFlag, *pFlag)
	toCsv(timesheet.getEntries(), timesheet.currMonth)

}
