package space

// Planet is the location where age will be calculated
type Planet string

// Possible planets
const (
	PlanetMercury Planet = "Mercury"
	PlanetVenus   Planet = "Venus"
	PlanetEarth   Planet = "Earth"
	PlanetMars    Planet = "Mars"
	PlanetJupiter Planet = "Jupiter"
	PlanetSaturn  Planet = "Saturn"
	PlanetUranus  Planet = "Uranus"
	PlanetNeptune Planet = "Neptune"
)

var periods = map[Planet]float64{
	PlanetMercury: 76005.3024,
	PlanetVenus:   194139.072,
	PlanetEarth:   315581.4976,
	PlanetMars:    593542.944,
	PlanetJupiter: 3743357.76,
	PlanetSaturn:  9295966.08,
	PlanetUranus:  26307031.65,
	PlanetNeptune: 52004185.92,
}

// Age calculates number of years in the given planet
func Age(seconds float64, planet Planet) float64 {
	period, ok := periods[planet]
	if !ok {
		period = 1
	}
	return float64(int64(seconds/period+0.5)) / 100
}
