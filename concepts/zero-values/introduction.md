# Introduction

In Go, every variable has a value.
Declaring a variable without assigning one sets it to the zero value for its type.

Basic types, including `bool`, numeric types, and `string`, each have a fixed zero value:

| Type                  | Zero Value |
| --------------------- | ---------- |
| bool                  | `false`    |
| int, int8, int64, etc. | `0`        |
| float32, float64      | `0`        |
| complex64, complex128 | `0+0i`     |
| string                | `""`       |

Pointers, functions, interfaces, slices, channels, and maps each have a zero value of `nil`.

Arrays and structs are never `nil` because they hold the data directly, not a reference to it.
Each element of an array and each field of a struct is initialized to its own type's zero value.
For example, `var a [3]int` creates the array `[0, 0, 0]`.

## Declaring Zero Values

Use `var` for a declaration:

```go
var myBool bool   // false
var mySlice []int // nil
```

Use a composite literal `T{}` for structs:

```go
type Person struct {
    Name string
    Age  int
}

myPerson := Person{} // Name: "", Age: 0
```
