package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var OPEN_WEATHER_API_KEY string

func main() {
	if err := SetEnv(); err != nil {
		log.Fatalf(err.Error())
	}

	http.HandleFunc("/", HandleWeather)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func SetEnv() error {
	if OPEN_WEATHER_API_KEY = os.Getenv("OPEN_WEATHER_API_KEY"); OPEN_WEATHER_API_KEY == "" {
		return fmt.Errorf("OPEN_WEATHER_API_KEY environment variable missing.")
	}
	return nil
}
