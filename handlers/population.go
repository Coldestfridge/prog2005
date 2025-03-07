package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// gets information about a specific country
func Getpopulation(w http.ResponseWriter, r *http.Request) {

	//splits the url into parts based on /
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 5 || pathParts[4] == "" {
		http.Error(w, "❌ Error: Missing country code. Use format /countryinfo/v1/population/{country_code}", http.StatusBadRequest)
		return
	}

	countryCode := strings.ToUpper(strings.TrimSpace(pathParts[4]))

	if len(countryCode) != 2 {
		http.Error(w, "❌ Error: Invalid country code. Use a two-letter ISO code (e.g., 'NO' for Norway).", http.StatusBadRequest)
		return
	}

	//adds the IP address and country code to URL
	apiURL := fmt.Sprintf("%salpha/%s", countries, countryCode)

	//connects to URL
	resp, err := http.Get(apiURL)
	if err != nil {
		http.Error(w, fmt.Sprintf("❌ Error: Failed to connect to API: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	//sends a statuscode if connection fails
	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("❌ Error: API returned status code %d. Country not found.", resp.StatusCode), http.StatusNotFound)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "❌ Error: Failed to read API response.", http.StatusInternalServerError)
		return
	}

	//retrives full name of country
	var countryData []struct {
		Name struct {
			Common string `json:"common"`
		} `json:"name"`
	}

	if err := json.Unmarshal(body, &countryData); err != nil || len(countryData) == 0 {
		http.Error(w, "❌ Error: Failed to parse country details response.", http.StatusInternalServerError)
		return
	}

	countryName := strings.ToLower(countryData[0].Name.Common)
	fmt.Println("🔍 Extracted Country Name:", countryName) //shows wich countryname is extracted

	// Extract and parse `limit=YYYY-YYYY`
	limit := r.URL.Query().Get("limit")
	var startYear, endYear int

	if limit != "" {
		yearRange := strings.Split(limit, "-")
		if len(yearRange) != 2 {
			http.Error(w, "❌ Error: Invalid limit format. Use YYYY-YYYY (e.g., 2010-2015).", http.StatusBadRequest)
			return
		}

		var err error
		startYear, err = strconv.Atoi(yearRange[0])
		if err != nil {
			http.Error(w, "❌ Error: Invalid start year in limit.", http.StatusBadRequest)
			return
		}

		endYear, err = strconv.Atoi(yearRange[1])
		if err != nil {
			http.Error(w, "❌ Error: Invalid end year in limit.", http.StatusBadRequest)
			return
		}
	}

	payload := map[string]string{"country": countryName}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "❌ Error: Failed to create JSON request payload.", http.StatusInternalServerError)
		return
	}

	populationURL := fmt.Sprintf("%scountries/population", countriesnow)

	//makes a post request
	req, err := http.NewRequest("POST", populationURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		http.Error(w, fmt.Sprintf("❌ Error: Failed to create request: %v", err), http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("❌ Error: Failed to connect to CountriesNow API: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	fmt.Println("🔍 CountriesNow API Response Status:", resp.Status)

	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("❌ Error: CountriesNow API returned status code %d. No population data found.", resp.StatusCode), http.StatusNotFound)
		return
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "❌ Error: Failed to read population API response.", http.StatusInternalServerError)
		return
	}

	fmt.Println("🔍 RAW POPULATION API RESPONSE:", string(body))

	var popData struct {
		Data struct {
			PopulationCounts []struct {
				Year  int `json:"year"`
				Value int `json:"value"`
			} `json:"populationCounts"`
		} `json:"data"`
	}

	if err := json.Unmarshal(body, &popData); err != nil {
		http.Error(w, fmt.Sprintf("❌ Error: Failed to parse population data: %v", err), http.StatusInternalServerError)
		return
	}

	if len(popData.Data.PopulationCounts) == 0 {
		http.Error(w, "❌ Error: No population data found for this country.", http.StatusNotFound)
		return
	}

	filteredPopulation := limitvalues{}

	for _, record := range popData.Data.PopulationCounts {
		if (startYear == 0 || record.Year >= startYear) && (endYear == 0 || record.Year <= endYear) {
			filteredPopulation = append(filteredPopulation, record)
		}
	}

	if len(filteredPopulation) == 0 {
		http.Error(w, "❌ Error: No population data found within the given year range.", http.StatusNotFound)
		return
	}

	total := 0
	for _, record := range filteredPopulation {
		total += record.Value
	}
	mean := total / len(filteredPopulation)

	filteredResponse := PopulationData{
		Mean:   mean,
		Values: filteredPopulation,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredResponse)
}
