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

You may have noticed struct types are absent from the above table. That is because the zero value for a struct type depends on its fields. Structs are set to their zero value when all of its fields are set to their respective zero value.

## Zero Value Construction

The `var` keyword can be used to construct any type to its zero value:

```go
var myBool bool
fmt.Printf("Zero value boolean: %#v", myBool)
// Output: Zero value boolean: false
```

```go
var mySlice []int
fmt.Printf("Zero value slice: %#v", mySlice)
// Output: Zero value slice: []int(nil)
```

When constructing the zero value for a struct type, all of the struct's fields will be set to their zero value:

```go
type Person struct {
  Name string
  Age  int
}

var myPerson Person
fmt.Printf("Zero value Person: %#v", myPerson)
// Output: Zero value Person: main.Person{Name:"", Age:0}
```

## Comparing with Nil

If you try to compare a type whose zero value is not `nil` to `nil`, your code will not compile. You will see a compiler error because booleans, numeric types, and strings can never be `nil` in Go.

```go
var myString string

if myString != nil { // invalid operation: myString != nil (mismatched types string and nil)
  fmt.Println("Do some work here.")
}
```

However, comparing a type whose zero value is `nil` to `nil` is acceptable:

```go
var mySlice []int

if mySlice != nil {
  fmt.Println("Do some work here.")
}
```
