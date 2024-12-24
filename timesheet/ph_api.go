package timesheet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type PublicHolidayAPI struct {
	apiBaseURL string
	logger     log.Logger
}

func NewPublicHolidayAPI(logger log.Logger) *PublicHolidayAPI {
	return &PublicHolidayAPI{
		apiBaseURL: "https://date.nager.at/api/v3/PublicHolidays/",
		logger:     logger,
	}
}

// Holidays models the response of https://date.nager.at/api/v2/PublicHolidays/2021/RO
type Holidays []struct {
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

// Fetch handles the low-level fetching on the API
func (api *PublicHolidayAPI) Fetch(year int) Holidays {
	client := &http.Client{}
	apiURL := api.apiBaseURL + fmt.Sprint(year) + "/RO"
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		level.Error(api.logger).Log("err", err.Error())
		os.Exit(1)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		level.Error(api.logger).Log("err", err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		level.Error(api.logger).Log("err", err.Error())
		os.Exit(1)
	}
	var responseObject Holidays
	json.Unmarshal(bodyBytes, &responseObject)

	return responseObject
}
