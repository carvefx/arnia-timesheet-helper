package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const apiBaseURL string = "https://date.nager.at/api/v2/PublicHolidays/"

// holidays models the response of https://date.nager.at/api/v2/PublicHolidays/2021/RO
type holidays []struct {
	Date        string      `json:"date"`
	LocalName   string      `json:"localName"`
	Name        string      `json:"name"`
	CountryCode string      `json:"countryCode"`
	Fixed       bool        `json:"fixed"`
	Global      bool        `json:"global"`
	Counties    interface{} `json:"counties"`
	LaunchYear  interface{} `json:"launchYear"`
	Type        string      `json:"type"`
}

// Process the data from the API into a slice of day-dates of public holidays
func getHolidaysByMonth(year int, month int) []int {
	h := fetchHolidays(year)
	var filteredHolidays []int = make([]int, 0)
	for _, day := range h {
		currDate, err := time.Parse("2006-01-02", day.Date)

		if err != nil {
			log.Fatal("Error parsing dates")
		}

		if int(currDate.Month()) == month {
			filteredHolidays = append(filteredHolidays, currDate.Day())
		}
	}
	return filteredHolidays
}

// handle the low-level fetching on the API
func fetchHolidays(year int) holidays {
	client := &http.Client{}
	apiURL := apiBaseURL + fmt.Sprint(year) + "/RO"
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var responseObject holidays
	json.Unmarshal(bodyBytes, &responseObject)

	return responseObject
}
