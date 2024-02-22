package internal

import (
	"fmt"
	"log"
	"net/http"
)

// NewWeatherHandler creates a handler that parses a request for weather data and returns the weather for
// a given location using the OpenWeather API.
func NewWeatherHandler(apiKey string, home *Location) http.HandlerFunc {
	client := NewWeatherClient(apiKey, "metric")

	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("Received request for weather data.")

		lat := r.URL.Query().Get("lat")
		long := r.URL.Query().Get("lon")

		var location *Location
		if lat == "" || long == "" {
			log.Print("Location parameters not supplied, defaulting to home")
			if home == nil {
				log.Print("Home location not set")
				return
			}
			location = home
		} else {
			log.Printf("Location received %s %s", lat, long)
			location = NewLocation(lat, long)
		}

		report, err := client.FetchWeatherReport(location.Lat, location.Long)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, report)
	}
}
