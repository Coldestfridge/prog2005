package handlers

// CountryInf holds general information about a country
type CountryInf struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flag       string            `json:"flag"`
	Flags      struct {
		PNG string `json:"png"`
	} `json:"flags"`
	Capital []string `json:"capital"`
	Cities  []string `json:"cities"`
}

type CountryResponse struct {
	Name       string            `json:"name"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flag       string            `json:"flag"`    // ✅ Using PNG URL instead of emoji
	Capital    string            `json:"capital"` // ✅ Single string, not array
	Cities     []string          `json:"cities"`
}

// PopulationData represents the historical population data
type PopulationData struct {
	Mean   int `json:"mean"`
	Values []struct {
		Year  int `json:"year"`
		Value int `json:"value"`
	} `json:"data"`
}

type limitvalues []struct {
	Year  int `json:"year"`
	Value int `json:"value"`
}

// StatusResponse struct for JSON response
type StatusResponse struct {
	CountriesNowAPI  int    `json:"countriesnowapi"`
	RestCountriesAPI int    `json:"restcountriesapi"`
	Version          string `json:"version"`
	Uptime           int64  `json:"uptime"` // in seconds
}
