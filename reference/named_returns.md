# Named Returns

In Go a return value can be named.

A good name for a return value can help its user to understand. Without looking at the code the following
function signature tells a lot more than `getPosition() (float64, float64, error)` would:

```go
func getPosition() (lat, long float64, err error) {
    // lat, long, err already exist
}
```

The second effect of using names for return values is, that these variables are already declared in the function.

With `named returns` also `naked returns` can be used. Although there are 3 return values, the `return` statement can stand alone:

```go
func getPosition() (lat, long float64, err error) {
    lat = 4.534
    long = 3.5273

    return
}
```

**Warning**: Because `naked returns` make functions harder to read (especially if they are longer than a few lines),
they are considered bad practice by the Go community.
