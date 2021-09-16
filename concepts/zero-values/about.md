# About

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

## Zero Value Construction

The `var` keyword can be used to construct any type to its zero value:

```go
func main() {
  var myBool bool
  fmt.Println("Zero value boolean:", myBool)
}
```

```go
func main() {
  var mySlice []int
  fmt.Println("Zero value slice:", mySlice)
}
```

When constructing the zero value for a struct type, all of the struct's fields will be set to their zero value:

```go
type Person struct {
  Name string
  Age  int
}

func main() {
  var myPerson Person
  fmt.Println("Zero value Person:", myPerson)
}
```

## Related Concepts

- [concept:go/nil]()
