package meteorology

import (
	"testing"
)

func TestTemperatureUnitString(t *testing.T) {
	tests := []struct {
		name string
		unit TemperatureUnit
		want string
	}{
		{
			name: "Celsius unit",
			unit: Celsius,
			want: "°C",
		},
		{
			name: "Fahrenheit unit",
			unit: Fahrenheit,
			want: "°F",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.unit.String(); got != tt.want {
				t.Errorf("unit.String()=%q, want %q", got, tt.want)
			}
		})
	}
}

func TestTemperatureString(t *testing.T) {
	tests := []struct {
		name string
		temp Temperature
		want string
	}{
		{
			name: "21 degree Celsius",
			temp: Temperature{21, Celsius},
			want: "21 °C",
		},
		{
			name: "32 degree Fahrenheit",
			temp: Temperature{32, Fahrenheit},
			want: "32 °F",
		},
		{
			name: "minus 17 degree Celsius",
			temp: Temperature{-17, Celsius},
			want: "-17 °C",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.temp.String(); got != tt.want {
				t.Errorf("temp.String()=%q, want %q", got, tt.want)
			}
		})
	}
}

func TestSpeedUnitString(t *testing.T) {
	tests := []struct {
		name string
		unit SpeedUnit
		want string
	}{
		{
			name: "kmPerHour",
			unit: KmPerHour,
			want: "km/h",
		},
		{
			name: "milesPerHour",
			unit: MilesPerHour,
			want: "mph",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.unit.String(); got != tt.want {
				t.Errorf("unit.String()=%q, want %q", got, tt.want)
			}
		})
	}
}

func TestSpeedString(t *testing.T) {
	tests := []struct {
		name  string
		speed Speed
		want  string
	}{
		{
			name:  "18 kilometers per hour",
			speed: Speed{18, KmPerHour},
			want:  "18 km/h",
		},
		{
			name:  "30 miles per hour",
			speed: Speed{30, MilesPerHour},
			want:  "30 mph",
		},
		{
			name:  "0 kilometers per hour",
			speed: Speed{0, KmPerHour},
			want:  "0 km/h",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.speed.String(); got != tt.want {
				t.Errorf("speed.String()=%q, want %q", got, tt.want)
			}
		})
	}
}

func TestMeteorologyDataString(t *testing.T) {
	tests := []struct {
		name            string
		meteorologyData MeteorologyData
		want            string
	}{
		{
			name:            "Athens",
			meteorologyData: MeteorologyData{"Athens", Temperature{21, Celsius}, "N", Speed{16, KmPerHour}, 63},
			want:            "Athens: 21 °C, Wind N at 16 km/h, 63% Humidity",
		},
		{
			name:            "Delhi",
			meteorologyData: MeteorologyData{"Delhi", Temperature{33, Celsius}, "W", Speed{2, MilesPerHour}, 23},
			want:            "Delhi: 33 °C, Wind W at 2 mph, 23% Humidity",
		},
		{
			name:            "San Francisco",
			meteorologyData: MeteorologyData{"San Francisco", Temperature{57, Fahrenheit}, "NW", Speed{19, MilesPerHour}, 60},
			want:            "San Francisco: 57 °F, Wind NW at 19 mph, 60% Humidity",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.meteorologyData.String(); got != tt.want {
				t.Errorf("meteorologyData.String()=%q, want %q", got, tt.want)
			}
		})
	}
}
