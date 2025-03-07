package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// GetCountry fetches general country info
func Getcountry(w http.ResponseWriter, r *http.Request) {

	// Extract country code from the request URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 5 || pathParts[4] == "" {
		http.Error(w, "❌ Error: Missing country code. Use format /countryinfo/v1/info/{country_code}", http.StatusBadRequest)
		return
	}
	countryCode := strings.ToUpper(strings.TrimSpace(pathParts[4]))

	// Validate country code format
	if len(countryCode) != 2 {
		http.Error(w, "❌ Error: Invalid country code. Use two-letter ISO code (e.g., 'NO' for Norway).", http.StatusBadRequest)
		return
	}

	// Extract limit from query parameters (default: 10)
	limit := 10 // Default limit
	queryParams := r.URL.Query()
	if limitParam, ok := queryParams["limit"]; ok {
		parsedLimit, err := strconv.Atoi(limitParam[0])
		if err != nil || parsedLimit <= 0 { // ✅ Prevents limit=0
			http.Error(w, "❌ Error: Limit must be a positive integer greater than 0.", http.StatusBadRequest)
			return
		}
		limit = parsedLimit
	}

	// Construct API URL
	apiURL := fmt.Sprintf("%salpha/%s", countries, countryCode)

	// Make API request
	resp, err := http.Get(apiURL)
	if err != nil {
		http.Error(w, fmt.Sprintf("❌ Error: Failed to connect to API: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Handle HTTP errors from external API
	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("❌ Error: API returned status code %d. Country not found.", resp.StatusCode), http.StatusNotFound)
		return
	}

	// Read API response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "❌ Error: Failed to read API response.", http.StatusInternalServerError)
		return
	}

	// Parse JSON response
	var countryData []CountryInf
	if err := json.Unmarshal(body, &countryData); err != nil {
		fmt.Println("❌ JSON Unmarshal Error:", err) // Print detailed error
		http.Error(w, fmt.Sprintf("❌ Error: Failed to parse JSON response: %v", err), http.StatusInternalServerError)
		return
	}

	// Ensure data is available
	if len(countryData) == 0 {
		http.Error(w, "❌ Error: No data found for this country code.", http.StatusNotFound)
		return
	}

	// Fetch cities using country name
	cities, err := FetchCities(countryData[0].Name.Common, limit)
	if err != nil {
		fmt.Println("❌ Error fetching cities:", err)
		cities = []string{"Unknown"} // Fallback value if API fails
	}

	// Convert `CountryInf` to `CountryResponse`
	response := ConvertToCountryResponse(countryData[0], cities)

	// Convert response to JSON and send to client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func FetchCities(countryName string, limit int) ([]string, error) {
	// Construct API URL for cities
	apiURL := fmt.Sprintf("%scountries/cities/q?country=%s", countriesnow, countryName)

	// Make API request
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("❌ Error: Failed to fetch cities: %v", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("❌ Error: Failed to read cities API response: %v", err)
	}

	// Define struct to match API response
	var result struct {
		Error bool     `json:"error"`
		Msg   string   `json:"msg"`
		Data  []string `json:"data"`
	}

	// Parse JSON
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("❌ Error: Failed to parse cities API response: %v", err)
	}

	// Check if the limit is smaller than the available number of cities
	if len(result.Data) > limit {
		result.Data = result.Data[:limit] // Trim list to requested limit
	}

	// Return city names
	return result.Data, nil
}

func ConvertToCountryResponse(country CountryInf, cities []string) CountryResponse {
	return CountryResponse{
		Name:       country.Name.Common,
		Continents: country.Continents,
		Population: country.Population,
		Languages:  country.Languages,
		Borders:    country.Borders,
		Flag:       country.Flags.PNG,             // Use PNG URL instead of emoji flag
		Capital:    FirstElement(country.Capital), // Convert to a single string
		Cities:     cities,
	}
}

func FirstElement(list []string) string {
	if len(list) > 0 {
		return list[0]
	}
	return ""
}
