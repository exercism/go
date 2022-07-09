package weatherforecast

import "fmt"

// Forecast will return predict for a city's weather.
func Forecast(city string) string {
	return fmt.Sprintf("%s will have a nice sunny day tomorrow!", city)
}
