# Hints

## 1. Implement the `Stringer` interface for type `TemperatureUnit`

- The type `TemperatureUnit` is already created for you and represents a temperature unit
- Note that there are already 2 constants created for you that are of this type: `Celsius` and `Fahrenheit`, representing a temperature in Celsius and Fahrenheit, respectively
- Add a `String()` method to the `TemperatureUnit` type so it satisfies the `Stringer` interface. This method must return the string `"°C"` if the temperature unit is Celsius or `"°F"` if the temperature unit is Fahrenheit

## 2. Implement the `Stringer` interface for type  `Temperature` 

 - Add a `String()` method to the `Temperature` type so it satisfies the `Stringer` interface
 - The `String()` method should return a string with the numeric value for the temperature and the temperature unit separated by a space (`<temperature> <unit>`)
 - The `Temperature` struct contains a `TemperatureUnit` and a `int`. You can use [fmt.Sprintf][sprintf] to help you format this string
 - Since `TemperatureUnit` already implements the `Stringer` interface (from task 1), the functions of the `fmt` package like [fmt.Sprintf][sprintf] will know how to format it when you use the `%v` or `%s` [formatting verbs][fmt]

## 3. Implement the `Stringer` interface for type `SpeedUnit`

- The type `SpeedUnit` is already created for you and represents a temperature unit
- Note that there are already 2 constants created for you that are of this type: `KmPerHour` and `MilesPerHour`, representing a speed in kilometers per hour and miles per hour, respectively
- Add a `String()` method to the `SpeedUnit` type so it satisfies the `Stringer` interface. This method must return the string `"km/h"` if the speed unit is kilometers per hour or `"mph"` if the speed unit is miles per hour.

## 4. Implement the `Stringer` interface for `Speed`

 - Add a `String()` method to the `Speed` type so it satisfies the `Stringer` interface
 - The `String()` method should return a string with the numeric value for the speed and the speed unit separated by a space (`<speed> <unit>`)
 - The `Speed` contains a `SpeedUnit` and a `int`. You can use [fmt.Sprintf][sprintf] to help you format this string
 - Since `TemperatureUnit` struct already implements the `Stringer` interface (from task 3), the functions of the `fmt` package like [fmt.Sprintf][sprintf] will know how to format it when you use the `%v` or `%s` [formatting verbs][fmt]
- To insert a `%` in the final string when using [fmt.Sprintf][sprintf], use `%%` in the formatting string.

## 5. Implement the `Stringer` interface for type `MeteorologyData`

 - The `Speed` contains a `Temperature` and a `Speed` fields, among other fields. You can use [fmt.Sprintf][sprintf] to help you format this string
 - The `String` method should return the meteorology data in the following format: `"<location>: <temperature>, Wind <wind_speed> at <wind_direction>, <humidity>% Humidity"`
 - Since the `Temperature` and `Speed` types already implement the `Stringer` interface (from task 2 and 4), the functions of the `fmt` package like [fmt.Sprintf][sprintf] will know how to format it when you use the `%v` or `%s` [formatting verbs][fmt]


[fmt]: https://pkg.go.dev/fmt
[sprint]: https://pkg.go.dev/fmt#Sprint
[sprintf]: https://pkg.go.dev/fmt#Sprintf
[yourbasic-enum]: https://yourbasic.org/golang/iota/#complete-enum-type-with-strings-best-practice
