# Instructions

Your team is working on a meteorology application. 
They have defined an API with various types and constants representing meteorological data, see file `meteorology.go`. 
  
Your task is to add suitable `String` methods to all types so that they implement interface `Stringer`. 

## 1. Implement the `Stringer` interface for type `TemperatureUnit`

After some discussion, the team have agreed that the unit of temperature will be either `Celsius` or `Fahrenheit`. Values should be formatted as shown in the examples below.

Make the `TemperatureUnit` type implement the `Stringer` interface by adding a `String` method to it. This method must return the string `"°C"` if the temperature unit is Celsius or `"°F"` if the temperature unit is Fahrenheit.

```go 
temperatureUnit := Celsius
celsiusUnit := Celsius
fahrenheitUnit := Fahrenheit

celsiusUnit.String()
// => °C
fahrenheitUnit.String()
// => °F
fmt.Sprint(celsiusUnit)
// => °C
```

## 2. Implement the `Stringer` interface for type  `Temperature`

Temperature values consist of an integer and a temperature unit. They should be formatted as in the examples below.

For that to happen, make the `Temperature` type implement the `Stringer` interface by adding a `String` method to it. This method should return a string with the numeric value for the temperature and the temperature unit separated by a space: `<temperature> <unit>`:


```go
celsiusTemp := Temperature{
    degree: 21,
    unit: Celsius,
}
celsiusTemp.String()
// => 21 °C
fmt.Sprint(celsiusTemp)
// => 21 °C

fahrenheitTemp := Temperature{
    degree: 75,
    unit: Fahrenheit,
}
fahrenheitTemp.String()
// => 75 °F
fmt.Sprint(fahrenheitTemp) 
// => 75 °F
```

## 3. Implement the `Stringer` interface for type `SpeedUnit`

After lengthy discussions, the team has agreed that the unit of wind speed will be either `KmPerHour` or `MilesPerHour`. Values should be formatted as the examples below.

For that to happen, make the `SpeedUnit` type implement the `Stringer` interface by adding a `String` method to it. This method must return the string `"km/h"` if the speed unit is kilometers per hour or `"mph"` if the speed unit is miles per hour:


```go 
mphUnit := MilesPerHour
mphUnit.String()
// => mph
fmt.Sprint(mphUnit)
// => mph

kmhUnit := KmPerHour
kmhUnit.String()
// => km/h
fmt.Sprint(kmhUnit)
// => km/h
```

## 4. Implement the `Stringer` interface for `Speed` 

Wind speed values consist of an integer and a speed unit. They should be formatted as in the example below.

For that to happen, make the `Speed` type implement the `Stringer` interface by adding a `String` method to it. This method should return a string with the numeric value for the speed and the speed unit separated by a space: `<speed> <unit>`:

```go 
windSpeedNow := Speed{
    magnitude: 18,
    unit: KmPerHour,
}
windSpeedNow.String(windSpeedNow)
// => 18 km/h
fmt.Sprintf(windSpeedNow)
// => 18 km/h

windSpeedYesterday := Speed{
    magnitude: 22,
    unit: MilesPerHour,
}
windSpeedYesterday.String(windSpeedYesterday)
// => 22 mph
fmt.Sprint(windSpeedYesterday)
// => 22 mph
```

## 5. Implement the `Stringer` interface for type `MetData`

Meteorological data specifies location, temperature, wind direction, wind speed
and humidity. It should be formatted as in the example below:

For that to happen, make the `MeteorologyData` type implement the `Stringer` interface by adding a `String` method to it. This method should return the meteorology data in the following format: `"<location>: <temperature>, Wind <wind_direction> at <wind_speed>, <humidity>% Humidity"`:

```go 
sfData := MeteorologyData{
    location: "San Francisco",
    temperature: Temperature{
        degree: 57,
        unit: Fahrenheit
    },
    windDirection: "NW",
    windSpeed: Speed{
        magnitude: 19,
        unit: MilesPerHour
    },
    humidity: 60
}

sfData.String()
// => San Francisco: 57 °F, Wind NW at 19 mph, 60% Humidity
fmt.Sprint(sfData) 
// => San Francisco: 57 °F, Wind NW at 19 mph, 60% Humidity
```
