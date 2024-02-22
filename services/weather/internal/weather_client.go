package internal

import (
	"fmt"
	"io"
	"net/http"
)

type WeatherClient struct {
	apiKey string
	units  string
}

func NewWeatherClient(apiKey string, units string) *WeatherClient {
	return &WeatherClient{
		apiKey: apiKey,
		units:  units,
	}
}

func (c *WeatherClient) FetchWeatherReport(lat, long string) (string, error) {
	resp, err := http.Get(
		fmt.Sprintf(
			"https://api.openweathermap.org/data/3.0/onecall?lat=%s&lon=%s&appid=%s",
			lat, long, c.apiKey,
		),
	)
	if err != nil {
		return "", err
	}
	weather, err := io.ReadAll(resp.Body)
	return string(weather), err
}
