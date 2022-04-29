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
	return [...]string{"C", "F"}[unit]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (t Temperature) String() string {
	return fmt.Sprintf("%d °%v", t.degree, t.unit)
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

type MetData struct {
	location  string
	temp      Temperature
	windDir   string
	windSpeed Speed
	humidity  int
}

func (m MetData) String() string {
	return fmt.Sprintf("%v: %v, Wind %v at %v, %v%% Humidity",
		m.location, m.temp, m.windDir, m.windSpeed, m.humidity)
}
