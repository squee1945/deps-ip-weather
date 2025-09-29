package ipweather

import "testing"

func TestGetWeather(t *testing.T) {
	ipw, err := New()
	if err != nil {
		t.Fatalf("New() returned an unexpected error: %v", err)
	}
	if _, err := ipw.GetWeather("bad ip"); err == nil {
		t.Errorf("Missing expected error")
	}
}
