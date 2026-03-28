# About

Go functions and methods can return multiple values:

```go
func GetCard() (Card, error) { ... }
```

Every return value must be assigned in order and separated by commas:

```go
card, err := GetCard()
```

## The Blank Identifier

Every position in the assignment declares a variable.
Declaring a variable without using it is not allowed in Go.
Use the blank identifier `_` in place of a variable name to skip that return value.

```go
card, _ := GetCard() // error intentionally ignored
_, err := GetCard()  // card ignored, only checking for error
```

## Error Handling

`error` is a built-in interface.
A `nil` error means no error occurred; any non-nil value means something went wrong.
By convention, functions that can fail should return an `error` as their last value.
The caller must check it before using any other return values:

```go
card, err := GetCard()
if err != nil {
    // unsafe to use card here
    return err
}
// safe to use card here
```

## The Comma-Ok Idiom

The two-value assignment pattern also appears outside of `error` handling.
Map lookups, type assertions, and channel receives all use a second boolean value, labeled `ok` by convention, to indicate success or the presence of data:

```go
m := map[string]int{"a": 1}
v, ok := m["a"] // ok is true
v, ok = m["z"]  // ok is false
```

```go
var i interface{} = "hello"
v, ok := i.(string) // ok is true
v, ok = i.(int)     // ok is false
```

```go
v, ok := <-ch // ok is false if ch is closed and empty
```

## Named Return Values

Return values can also be given names in the function signature.
These variables are then accessible within the body.

```go
func divide(a, b int) (quotient, remainder int) {
    quotient = a / b
    remainder = a % b
    return quotient, remainder
}
```

Named return values document the meaning of each return position.
They can also be used with a naked return, one written without values.
Naked returns implicitly return the current values of the named variables:

```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}
```

In longer functions, naked returns often hurt readability because it is unclear what's being returned.
Prefer explicit return values unless the function is short enough that the signature is always in view.

To learn more, check [Go by Example: Multiple Return Values][gobyexample], [Effective Go: Multiple Returns][effective_go], or the [Tour of Go][tour].

[tour]: https://go.dev/tour/basics/6
[gobyexample]: https://gobyexample.com/multiple-return-values
[effective_go]: https://go.dev/doc/effective_go#multiple-returns
