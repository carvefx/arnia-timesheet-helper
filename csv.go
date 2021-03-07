package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

var fileHeader = []string{"date", "account_id/id", "journal_id/id", "name", "unit_amount"}

func toCsv(entries []timesheetDay, month int) {
	currDate := time.Now()
	fileSuffix := fmt.Sprintf("%d-%d", currDate.Year(), month)
	file, err := os.Create("timesheet-" + fileSuffix + ".csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(fileHeader)
	checkError("Cannot write to file", err)

	for _, entry := range entries {
		err := writer.Write(parseEntryIntoRow(entry))
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func parseEntryIntoRow(entry timesheetDay) []string {
	row := make([]string, 5)

	row[0] = entry.date.Format("2006-01-02")
	row[1] = "__export__.account_analytic_account_" + fmt.Sprint(entry.projectID)
	row[2] = "hr_timesheet.analytic_journal"
	row[3] = entry.status.description
	row[4] = "8"

	return row
}
