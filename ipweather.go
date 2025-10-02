package ipweather

import (
	"fmt"
	"net"
)

type IPWeather struct{}

func New() (*IPWeather, error) {
	return &IPWeather{}, nil
}

type WeatherDetails struct {
	City        string
	Region      string
	Country     string
	Temperature float64
	Conditions  string
	Humidity    int
}

func (iw *IPWeather) GetWeather(ip string) (*WeatherDetails, error) {
	if net.ParseIP(ip) == nil {
		return nil, fmt.Errorf("invalid IP address %q", ip)
	}
	return &WeatherDetails{
		City:        "Example City",
		Region:      "Example Region",
		Country:     "Example Country",
		Temperature: 25.5,
		Conditions:  "Sunny and mild",
		Humidity:    60,
	}, nil
}
