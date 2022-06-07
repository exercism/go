# Introduction

In Go, a `struct` is a sequence of named elements called _fields_, each field has a name and type. The name of a field must be unique within the struct.
Structs can be compared with _classes_ in the Object-Oriented Programming paradigm.

You create a new struct by using the `type` keyword and the **_built-in type_** `struct`, and explicitly define the name and type of the fields.
For example, to define a `Shape` struct:

```go
type Shape struct{
    name string
    size int
}
```

Once you have defined the `struct`, you need to create a new instance defining the fields using their field name
in any order:

```go
s := Shape{
    name: "Square",
    size: 25,
}
```

To read or modify instance fields, use the `.` notation:

```go
// Update the name and the size
s.name = "Circle"
s.size = 35
fmt.Printf("name: %s size: %d\n", s.name, s.size)
// Output: name: Circle size: 35
```

Sometimes it can be nice to have a _constructor_, which is a function to create an instance of a struct.
The following constructor can be used to create a new instance of `Shape` and automatically set a default-value for the size of the shape:

```go
func NewShape(name string) Shape {
	return Shape{
		name: name,
		size: 100, //default-value for size is 100
	}
}
```

Using a constructor can have the following advantages:
* validation of the given values
* handling of default-values

Another advantage will be shown with an example. When importing the `package shapes` containing the
struct `Shape` the unexported fields `name` and `size` cannot be accessed. Using a constructor the initial value can nevertheless be set:
```go
rect := shapes.NewShape("Rectangle")
```
