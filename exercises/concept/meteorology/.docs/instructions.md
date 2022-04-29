# Instructions

Your team is working on a meteorology application. 
They have defined an API with various types and constants representing meteorological data, see file `meteorology.go`. 
  
Your task is to add suitable `String` methods to all types so that they implement interface `Stringer`. 

Values should be formatted as in the examples below: 

```go 
temperatureUnit := Celsius
fmt.Println(temperatureUnit)
// => C 

temperature := Temperature{21, Celsius}
fmt.Println(temperature) 
// => 21 °C

speedUnit := MilesPerHour
fmt.Println(speedUnit)
// => mph

windSpeed := Speed{18, KmPerHour}
fmt.Println(windSpeed)
// => 18 km/h

metData := MetData{"San Francisco", Temperature{57, Fahrenheit}, "NW", Speed{19, MilesPerHour}, 60},
fmt.Println(metData) 
// => San Francisco: 57 °F, Wind NW at 19 mph, 60% Humidity
```

In the test data, you can assume that

- the unit of temperature will be either `Celsius` or `Fahrenheit` 
- the unit of wind speed will be either `KmPerHour` or `MilesPerHour` 
