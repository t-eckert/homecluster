package main

import (
	"log"
	"net/http"

	"github.com/t-eckert/homecluster/services/weather/internal"
)

func main() {
	config := internal.NewConfigFromEnv().WithDefaults()

	http.HandleFunc("/", internal.NewWeatherHandler(config.ApiKey, config.Home))

	log.Printf("Serving on http://localhost%s\n", config.Port)
	if config.Home == nil {
		log.Print("Home location not set, no default location will be used.")
	} else if config.Home.Lat != "" && config.Home.Long != "" {
		log.Printf("Default location set to %s %s", config.Home.Lat, config.Home.Long)
	}

	log.Fatal(http.ListenAndServe(config.Port, nil))
}
