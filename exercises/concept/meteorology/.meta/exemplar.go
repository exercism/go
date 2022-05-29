package meteorology

import (
	"fmt"
)

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

func (unit TemperatureUnit) String() string {
	return [...]string{"°C", "°F"}[unit]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (t Temperature) String() string {
	return fmt.Sprintf("%d %v", t.degree, t.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

func (unit SpeedUnit) String() string {
	return [...]string{"km/h", "mph"}[unit]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (sp Speed) String() string {
	return fmt.Sprintf("%d %v", sp.magnitude, sp.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

func (m MeteorologyData) String() string {
	return fmt.Sprintf("%v: %v, Wind %v at %v, %v%% Humidity",
		m.location, m.temperature, m.windDirection, m.windSpeed, m.humidity)
}
