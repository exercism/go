# About

In Go, every variable has a value.
Declaring a variable without assigning a value sets it to the zero value for its type.

Basic types, including `bool`, numeric types, and `string`, each have a fixed zero value:

| Type                   | Zero Value |
| ---------------------- | ---------- |
| bool                   | `false`    |
| int, int8, int64, etc. | `0`        |
| float32, float64       | `0`        |
| complex64, complex128  | `0+0i`     |
| string                 | `""`       |

Types that reference underlying data, including pointers, functions, interfaces, slices, channels, and maps, each have a zero value of `nil`.
However, arrays and structs are never `nil` because they hold the data directly, not a reference to it.
Each element of an array and each field of a struct is initialized to its own type's zero value.
For example, `var a [3]int` creates the array `[0, 0, 0]`.

## Declaring Zero Values

`var` without an initial value gives the zero value for any type:

```go
var myBool bool   // false
var mySlice []int // nil
var p Person      // Person{Name: "", Age: 0}
```

For structs, a composite literal `T{}` is an alternative:

```go
myPerson := Person{} // equivalent to var myPerson Person
```

`new(T)` returns a pointer to the zero value of any type.
It is most commonly used with structs, though it works for basic types too:

```go
p := new(Person) // *Person, pointing to Person{Name: "", Age: 0}
s := new(string) // *string, pointing to ""
i := new(int)    // *int, pointing to 0
```

## Why Zero Values Matter

In Go, a zero value represents a natural and useful starting state.

A `bool` defaults to `false`, which can work as a flag:

```go
var done bool // action not done yet
done = true   // action now done
```

An integer defaults to `0`, which can work as a counter:

```go
var count int
count++ // 1
count++ // 2
```

For structs, each field starts at its own type's zero value.
A `Stack` with a single `[]string` field is already a working empty stack when declared.
Because `append` allocates a new backing array when called on a nil slice, no constructor is needed:

```go
type Stack struct {
    items []string
}

func (s *Stack) Push(v string) {
    s.items = append(s.items, v)
}

func (s *Stack) IsEmpty() bool {
    return len(s.items) == 0
}

var s Stack
fmt.Println(s.IsEmpty()) // true
s.Push("a") // append allocates a new backing array
fmt.Println(s.IsEmpty()) // false
```

## Working with Nil

Most types with `nil` zero values require initialization before use or they panic.
A `nil` slice is the exception: you can range over it, append to it, and pass it to `len` and `cap` safely.

A map should be initialized before storing a value:

```go
var m map[string]int
m["key"] = 1 // panic
```

A map is commonly initialized two different ways:

```go
m1 := make(map[string]int)
m2 := map[string]int{}
m1["key"] = 1 // ok
m2["key"] = 1 // ok
```

Check a pointer for `nil` before using it:

```go
var p *int

if p == nil {
    fmt.Println("no value assigned")
    return
}

fmt.Println(*p) // p is not nil
```

Basic types always hold a concrete value, so comparing them to `nil` is a compile error:

```go
var myString string

if myString == nil {
    // compile error: mismatched types
}
```
