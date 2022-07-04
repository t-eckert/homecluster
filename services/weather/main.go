package main

import (
	"log"
	"net/http"
	"os"
)

var ApiKey string
var DefaultLocation *Location

func main() {
	// Get configuration from the environment
	if ApiKey = os.Getenv("OPEN_WEATHER_API_KEY"); ApiKey == "" {
		log.Fatalf("OPEN_WEATHER_API_KEY environment variable missing.")
	}

	lon, lat := os.Getenv("LONGITUDE"), os.Getenv("LATITUDE")
	if lon != "" && lat != "" {
		DefaultLocation = &Location{lon, lat}
	}

	http.HandleFunc("/", HandleWeather)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
