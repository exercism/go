# Introduction

Go does not have a concept of empty, null, or undefined for variable values. Variables declared without an explicit initial value default to the zero value for their respective type.

The zero value for primative types such as booleans, numeric types, and strings are `false`, `0`, and `""`, respectively.

The identifier `nil`, meaning zero, is the zero value for more complex types such as pointers, functions, interfaces, slices, channels, and maps.

The following table details the zero value for Go's types.

| Type      | Zero Value |
| --------- | ---------- |
| boolean   | `false`    |
| numeric   | `0`        |
| string    | `""`       |
| pointer   | `nil`      |
| function  | `nil`      |
| interface | `nil`      |
| slice     | `nil`      |
| channel   | `nil`      |
| map       | `nil`      |

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
