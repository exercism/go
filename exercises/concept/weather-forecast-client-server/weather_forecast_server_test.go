package weatherforecast_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

const addr = "127.0.0.1:9898"

func TestForecastServer(t *testing.T) {
	testCases := []struct {
		method   string
		path     string
		body     string
		response string
		code     int
	}{
		{http.MethodGet, fmt.Sprintf("%s/", addr), "", "Please use '/city' to ask for weather status.", http.StatusBadRequest},
		{http.MethodGet, fmt.Sprintf("%s/City", addr), "", "Please use '/city' to ask for weather status.", http.StatusBadRequest},
		{http.MethodGet, fmt.Sprintf("%s/city", addr), "", "Golinocus will have a nice sunny day tomorrow!", http.StatusOK},
		{http.MethodPost, fmt.Sprintf("%s/city", addr), "Logob", "Logob", http.StatusOK},
		{http.MethodPost, fmt.Sprintf("%s/city", addr), "Logog", "city name incorrect.", http.StatusBadRequest},
		{http.MethodPut, fmt.Sprintf("%s/city", addr), "logob", "", http.StatusMethodNotAllowed},
	}
	for i, c := range testCases {
		srv := ForecastServer(addr, 5*time.Second)
		go func() {
			err := srv.ListenAndServe()
			if err != http.ErrServerClosed {
				t.Fatal(err)
			}
		}()
		client := http.Client{}
		req, err := http.NewRequest(c.method, c.path, strings.NewReader(c.body))
		if err != nil {
			t.Fatalf("could not create request. error: %s", err)
		}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("could not make the request. error: %s", err)
		}
		if c.code != resp.StatusCode {
			t.Fatalf("expected response code %d, received %d. test case number: %d", c.code, resp.StatusCode, i+1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		_ = resp.Body.Close()
		if string(b) != c.response {
			t.Fatalf("expected response: %s, received %s. test case number: %d", c.response, string(b), i+1)
		}
		srv.Close()
	}
}
