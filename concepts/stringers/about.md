# About

[Stringer][stringer-interface] is an interface for defining the string format of values.

The interface consists of a single `String` method:
 
```go
type Stringer interface {
    String() string
}
```
 
Types that want to implement this interface must have a `String()` method that returns a human-friendly string representation of the type. The [fmt][fmt-package] package (and many others) will look for this method to print values.

## Example: Distances

Assume we are working on an application that deals with geographical distances measured in different units. 
We have defined types `DistanceUnit` and `Distance` as follows: 
 
```go 
type DistanceUnit int

const (
	Kilometer    DistanceUnit = 0
	Mile         DistanceUnit = 1
	NauticalMile DistanceUnit = 2
)
 
type Distance struct {
	number float64
	unit   DistanceUnit
} 
```

These types do not implement interface `Stringer` as they lack the `String` method.
Hence `fmt` functions will print `Distance` values using Go's "default format":

```go 
var distances = []Distance{
	{number: 790.7, unit: Kilometer},
	{number: 415.2, unit: Mile},
	{number: 10_500, unit: NauticalMile},
} 
fmt.Println(distances)
// Output: [{790.7 0} {415.2 1} {10500 2}]
```

In order to make the output more informative, we implement interface `Stringer` by adding a `String` method to each type:

```go
func (sc DistanceUnit) String() string {
	units := []string{"km", "mi", "nmi"}
	return units[sc]
}

func (d Distance) String() string {
	return fmt.Sprintf("%v %v", d.number, d.unit)
} 
```
 
`fmt` package functions will call these methods when formatting `Distance` values:

```go
fmt.Println(distances)
// Output: [790.7 km 415.2 mi 10500 nmi]
```

[stringer-interface]: https://pkg.go.dev/fmt#Stringer
[fmt-package]: https://pkg.go.dev/fmt
