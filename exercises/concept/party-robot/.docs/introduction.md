# Introduction

Go programs are organized into packages. 
A package is a collection of source files in the same directory that are compiled together. 
Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package. 
Go determines if an item can be called by code in other packages through how it is declared.
To make a function, type, variable, constant or struct field externally visible (known as `exported`) the name must start with a capital letter.

For example:

```go
package greeting
// Hello is a public function (callable from other packages)
func Hello(name string) string {
    return "Hello " + name
}
// hello is a private function (not callable from other packages)
func hello(name string) string {
    return "Hello " + name
}
```

Go provides an in-built package called `fmt` (format package) which offers a variety of functions to manipulate the format of input and output.
The most commonly used function is `Sprintf`, which uses verbs like `%s` to interpolate values into a string and returns that string.

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

`fmt` contains other functions for working with strings, such as `Println` which simply prints the arguments it receives to the console and `Printf` which formats the input in the same way as `Sprintf` before printing it. Note that `Println` has built-in whitespace and newline logic when printing output to the console, whereas `Printf` does not.
