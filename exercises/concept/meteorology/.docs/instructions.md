# Instructions

Your team is working on a meteorology application. 
They have defined an API with various types and constants representing meteorological data, see file `meteorology.go`. 
  
Your task is to add suitable `String` methods to all types so that they implement interface `Stringer`. 

## 1. Implement the `Stringer` interface for type `TemperatureUnit`

After some discussion, the team have agreed that the unit of temperature will be either `Celsius` or `Fahrenheit`. Values should be formatted as shown below: 

```go 
temperatureUnit := Celsius
fmt.Println(temperatureUnit)
// Output: C 
```

## 2. Implement the `Stringer` interface for type  `Temperature`

Temperature values consist of an integer and a temperature unit. They should be formatted as in the example below: 

```go
temperature := Temperature{21, Celsius}
fmt.Println(temperature) 
// Output: 21 °C
```

## 3. Implement the `Stringer` interface for type `SpeedUnit`

After lengthy discussions, thee team has agreed that the unit of wind speed will be either `KmPerHour` or `MilesPerHour`. Values should be formatted as shown below: 

```go 
speedUnit := MilesPerHour
fmt.Println(speedUnit)
// Output: mph
```

## 4. Implement the `Stringer` interface for `Speed` 

Wind speed values consist of an integer and a speed unit. They should be formatted as in the example below: 

```go 
windSpeed := Speed{18, KmPerHour}
fmt.Println(windSpeed)
// Output: 18 km/h
```

## 5. Implement the `Stringer` interface for type `MetData`

Meteorological data specifies location, temperature, wind direction, wind speed
and humidity. It should be formatted as in the example below: 

```go 
metData := MetData{"San Francisco", Temperature{57, Fahrenheit}, "NW", Speed{19, MilesPerHour}, 60}
fmt.Println(metData) 
// Output: San Francisco: 57 °F, Wind NW at 19 mph, 60% Humidity
```
