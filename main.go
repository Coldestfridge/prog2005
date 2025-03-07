package main

import (
	"prog2005/handlers"
	"log"
	"net/http"
	"os"
)

func main() {

	// Handle port assignment (either based on environment variable, or local override)
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}

	// handler endpoints
	http.HandleFunc("/countryinfo/v1/info/", handlers.Getcountry)
	http.HandleFunc("/countryinfo/v1/population/", handlers.Getpopulation)
	http.HandleFunc("/countryinfo/v1/status/", handlers.GetStatus)

	// Start server
	log.Println("Starting server on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
