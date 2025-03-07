# prog2005
# 🌍 Country Info API

This is a RESTful API written in **Go (Golang)** that provides information about countries, including population data and general country details. The API fetches data from external sources like **CountriesNow** and **RestCountries**.

---

## 🚀 Installation

1. **Install Go** if you haven't already:  
   👉 [Download Go](https://golang.org/dl/)

2. **Clone this repository:**
   ```sh
   git clone https://github.com/yourusername/countryinfo-api.git
   cd countryinfo-api


## Usage
1. Start the server
2. when the server is started you can get information by using one of these commands in your browser
   
   To get information about a specific country:
   http://localhost:8080/countryinfo/v1/info/X    (X should be replaced with country code)

   To get population information about a specific country:
   http://localhost:8080/countryinfo/v1/population/X?limit=2010-2015
   (limit can be changed to other years or be removed if you want alle tha data the API has)

   To get status information about the APIs:
   http://localhost:8080/countryinfo/v1/status/

## Project structure
countryinfo-api/
│-- main.go
│-- handlers/
│   │-- country.go
│   │-- population.go
│   │-- status.go
│-- README.md
