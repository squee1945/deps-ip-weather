package ipweather

import (
	"fmt"
	"net"
)

type IPWeather struct{}

func New() (*IPWeather, error) {
	return &IPWeather{}, nil
}

func (iw *IPWeather) GetWeather(ip string) (string, error) {
	if net.ParseIP(ip) == nil {
		return "", fmt.Errorf("invalid IP address %q", ip)
	}
	return "Sunny and mild", nil
}
