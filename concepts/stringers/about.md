# About

[Stringer][stringer-interface] is an interface for defining the string format of values.

The interface consists of a single `String` method:
 
```go
type Stringer interface {
    String() string
}
```
 
Types that want to implement this interface must have a `String()` method that returns a human-friendly string representation of the type. The [fmt][fmt-package] package (and many others) will look for this method to format and print values.

## Example: Distances

Assume we are working on an application that deals with geographical distances measured in different units. 
We have defined types `DistanceUnit` and `Distance` as follows: 
 
```go 
type DistanceUnit int

const (
	Kilometer    DistanceUnit = 0
	Mile         DistanceUnit = 1
)
 
type Distance struct {
	number float64
	unit   DistanceUnit
} 
```

In the example above, `Kilometer` and `Mile` ane constants of type `DistanceUnit`.

These types do not implement interface `Stringer` as they lack the `String` method.
Hence `fmt` functions will print `Distance` values using Go's "default format":

```go
mileUnit := Mile
fmt.Sprint(mileUnit)
// => 1
// The result is '1' because that is the underlying value of the 'Mile' contant (see contant declarations above) 

dist := Distance{number: 790.7, unit: Kilometer}
fmt.Sprint(dist)
// => {790.7 0}
// not a very useful output!
```

In order to make the output more informative, we implement interface `Stringer` for `DistanceUnit` and `Distance` types by adding a `String` method to each type:

```go
func (sc DistanceUnit) String() string {
	units := []string{"km", "mi"}
	return units[sc]
}

func (d Distance) String() string {
	return fmt.Sprintf("%v %v", d.number, d.unit)
}
```
 
`fmt` package functions will call these methods when formatting `Distance` values:

```go
kmUnit := Kilometer
kmUnit.String()
// => km

mileUnit := Mile
mileUnit.String()
// => mi

dist := Distance{
	number: 790.7,
	unit: Kilometer,
}
dist.String()
// => 790.7 km
```

[stringer-interface]: https://pkg.go.dev/fmt#Stringer
[fmt-package]: https://pkg.go.dev/fmt