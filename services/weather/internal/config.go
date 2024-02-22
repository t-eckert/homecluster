package internal

import (
	"log"
	"os"
)

type Config struct {
	ApiKey string
	Home   *Location
	Port   string
}

func NewConfig(apiKey string, home *Location, port string) *Config {
	return &Config{
		ApiKey: apiKey,
		Home:   home,
		Port:   port,
	}
}

func NewConfigFromEnv() *Config {
	apiKey := os.Getenv("OPEN_WEATHER_API_KEY")
	if apiKey == "" {
		log.Fatalf("OPEN_WEATHER_API_KEY is required")
	}

	home := &Location{
		Lat:  os.Getenv("HOME_LAT"),
		Long: os.Getenv("HOME_LON"),
	}

	port := os.Getenv("PORT")

	return NewConfig(apiKey, home, port)
}

func (c *Config) WithDefaults() *Config {
	if c.Port == "" {
		c.Port = ":8080"
	}

	return c
}
