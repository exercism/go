package weather

var (
	CurrentCondition string
	CurrentLocation  string
)

func Log(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
