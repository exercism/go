package weatherforecastclient

import "time"

type WeatherClient struct {
	// Complete the WeatherClient
}

func NewWeatherClient(URL string, timeout time.Duration) *WeatherClient {
	panic("Implement NewWeatherClient here...")
}

func (wc *WeatherClient) TodayWeatherClient(city string) (string, error) {
	panic("Implement todayWeatherClient here...")
}

func (wc *WeatherClient) FutureWeatherQuery(city string, futureDate int) (string, error) {
	panic("Implement futureWeatherQuery here...")
}
