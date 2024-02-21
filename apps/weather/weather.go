package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Location struct {
	latitude, longitude string
}

// HandleWeather parses a request for weather data and returns the weather for
// a given location using the OpenWeather API.
func HandleWeather(w http.ResponseWriter, r *http.Request) {
	location, err := parseLocation(r)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	weather, err := fetchWeather(location)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, weather)
}

func parseLocation(r *http.Request) (Location, error) {
	var loc Location

	loc.latitude = r.URL.Query().Get("lat")
	loc.longitude = r.URL.Query().Get("lon")

	if (loc.latitude == "" || loc.longitude == "") && DefaultLocation != nil {
		return *DefaultLocation, nil
	}

	if loc.latitude == "" {
		return loc, fmt.Errorf("Missing query param for latitude. Try ?lat=<latitude>.")
	}

	if loc.longitude == "" {
		return loc, fmt.Errorf("Missing query param for longitude. Try ?lon=<longitude>.")
	}

	return loc, nil
}

func fetchWeather(location Location) (string, error) {
	resp, err := http.Get(
		fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s&units=metric",
			location.latitude, location.longitude, ApiKey),
	)
	if err != nil {
		return "", err
	}

	weather, err := ioutil.ReadAll(resp.Body)
	return string(weather), err
}
