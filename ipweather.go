package ipweather

import (
	"fmt"
	"net"
)

type IPWeather struct {
	IP string
}

func New(ip string) (*IPWeather, error) {
	if net.ParseIP(ip) == nil {
		return nil, fmt.Errorf("invalid IP address %q", ip)
	}
	return &IPWeather{
		IP: ip,
	}, nil
}

func (iw *IPWeather) GetWeather() (string, error) {
	return "Sunny and mild", nil
}
