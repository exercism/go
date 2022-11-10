package weatherforecastclient

import (
	"net/http"
	"time"
)

type WeatherClient struct {
	// Complete the WeatherClient
}

func NewWeatherClient(URL string, timeout time.Duration) *WeatherClient {
	panic("Implement NewWeatherClient here...")
}

func (wc *WeatherClient) TodayWeatherClient(city string) (*http.Response, error) {
	panic("Implement todayWeatherClient here...")
}

func (wc *WeatherClient) FutureWeatherQuery(city string, futureDate int) (*http.Response, error) {
	panic("Implement futureWeatherQuery here...")
}

func (wc *WeatherClient) WeatherReader(resp *http.Response, err error) (string, error) {
	panic("Implement WeatherReader here...")
}
