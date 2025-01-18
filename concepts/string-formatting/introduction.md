# Introduction

Go provides an in-built package called `fmt` (format package) which offers a variety of functions to manipulate the format of input and output.
The most commonly used function is `Sprintf`, which uses _verbs_ like `%s` to interpolate values into a string and returns that string.

```go
import "fmt"

food := "taco"
fmt.Sprintf("Bring me a %s", food)
// Returns: Bring me a taco
```

In Go floating point values are conveniently formatted with Sprintf's verbs: `%g` (compact representation), `%e` (exponent) or `%f` (non exponent).
All three verbs allow the field's width and numeric position to be controlled.

```go
import "fmt"

number := 4.3242
fmt.Sprintf("%.2f", number)
// Returns: 4.32
```

You can find a full list of available verbs in the [format package documentation][fmt-docs].

`fmt` contains other functions for working with strings, such as `Println` which simply prints the arguments it receives to the console and `Printf` which formats the input in the same way as `Sprintf` before printing it.

[fmt-docs]: https://pkg.go.dev/fmt
