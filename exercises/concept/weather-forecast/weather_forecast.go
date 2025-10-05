package weather

var (
	CurrentCondition string
	CurrentLocation  string
)

func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
