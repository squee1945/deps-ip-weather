package ipweather

import "testing"

func TestNew(t *testing.T) {
	if _, err := New("bad ip"); err == nil {
		t.Errorf("Missing expected error")
	}
}
