package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// tracks time since start
var startTime = time.Now()

// GetStatus checks API availability and returns service status
func GetStatus(w http.ResponseWriter, r *http.Request) {

	// Check API statuses
	countriesNowStatus := checkAPI(countriesnow + "countries")
	restCountriesStatus := checkAPI(countries + "all")

	// Calculate uptime
	uptime := time.Since(startTime).Seconds()

	// Construct JSON response
	statusResponse := StatusResponse{
		CountriesNowAPI:  countriesNowStatus,
		RestCountriesAPI: restCountriesStatus,
		Version:          "v1",
		Uptime:           int64(uptime),
	}

	// Convert response to JSON and send to client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(statusResponse)
}

// checkAPI makes a GET request to the API and returns the status code
func checkAPI(apiURL string) int {
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("❌ Error checking API:", apiURL, err)
		return 0 // Return 0 if the API is unreachable
	}
	defer resp.Body.Close()

	return resp.StatusCode
}
