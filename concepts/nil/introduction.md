# Introduction

Go does not have a concept of undefined for variable values. Variables declared without an explicit initial value default to the [zero value][zero_values] for their respective type.

The zero value for primative types such as booleans, numeric types, and strings are `false`, `0`, and `""`, respectively.

The identifier `nil`, meaning zero, is the zero value for more complex types such as pointers, interfaces, channels, maps, slices, and functions.

The following table details the zero value for Go's types.

| Type      | Zero Value |
| --------- | ---------- |
| bool      | false      |
| numeric   | 0          |
| string    | ""         |
| pointer   | nil        |
| interface | nil        |
| channel   | nil        |
| map       | nil        |
| slice     | nil        |
| func      | nil        |

## Zero Value Construction

The `var` keyword can be used to construct any type to its zero value:

```go
func main() {
  var myBool bool
  fmt.Println(myBool)
}
```

```go
func main() {
  var myMap map[int]int
  fmt.Println(myMap)
}
```

## Comparing Types to Nil

Comparing a type whose zero value is not `nil` to `nil` is an error:

```go
func main() {
  var myString string

  if myString != nil { // invalid operation: myString != nil (mismatched types string and nil)
    fmt.Println(myString)
  }
}
```

However, comparing a type whose zero value is `nil` to `nil` is acceptable:

```go
func main() {
  var mySlice []int

  if mySlice != nil {
    fmt.Println(mySlice)
  }
}
```

[zero_values]: https://golang.org/ref/spec#The_zero_value
