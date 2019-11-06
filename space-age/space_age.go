package space

import "math"

const earthRevolution float64 = 31557600
type Planet string

func convertBy(planet Planet) func(ageInSeconds float64) float64 {
	Planets := map[Planet]float64{
		"Earth":   1,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	planetRevolution := earthRevolution * Planets[planet]

	return func(ageInSeconds float64) float64 {
		return math.Round(ageInSeconds / planetRevolution * 100) / 100
	}
}

func Age(ageInSeconds float64 , planet Planet) float64 {
	converter := convertBy(planet)
	return converter(ageInSeconds)
}