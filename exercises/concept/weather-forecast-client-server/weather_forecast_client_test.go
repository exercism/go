package weatherforecast_test

import (
	"net/http"
	"testing"
)

func TestForecastClient(t *testing.T) {
	testCases := []struct {
		city        string
		serverAddr  string
		method      string
		expected    string
		errExpected bool
	}{
		{"", addr, http.MethodPost, "", true},
		{"Logob", addr, http.MethodPut, "", true},
		{"", addr, http.MethodGet, "Golinocus will have a nice sunny day tomorrow!", false},
		{"Logob", addr, http.MethodPost, "Logob will have a nice sunny day tomorrow!", false},
		{"Logog", addr, http.MethodPost, "city name incorrect.", true},
	}
	for i, c := range testCases {
		resp, err := ForecastClient(c.city, c.serverAddr, c.method)
		if (err == nil) && c.errExpected {
			t.Fatalf("expected error, returned nil. test case number: %d", i+1)
		}
		if (err != nil) && !c.errExpected {
			t.Fatalf("did not expect an error. received: %s. test case number: %d", err, i+1)
		}
		if resp != c.expected {
			t.Fatalf("expected %s, received %s. test case number: %d", c.expected, resp, i+1)
		}
	}
}
