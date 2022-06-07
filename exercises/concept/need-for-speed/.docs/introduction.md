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
in any order. To initialize instances of structs Go offers two ways:

```go
s1 := Shape{
    name: "Square",
    size: 25,
}

s2 := new(Shape)
fmt.Printf("name: %s size: %d\n", s2.name, s2.size)
// Output: name:  size: 0
```

As shown in this example `func new(Type)` can be used to create an newly created zero value of the given type.

To read or modify instance fields, use the `.` notation:

```go
// Update the name and the size
s2.name = "Circle"
s2.size = 35
fmt.Printf("name: %s size: %d\n", s2.name, s2.size)
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
struct `Shape` the unexported fields `name` and `size` cannot be accessed. Using a constructor the initial value
can be set:
```go
rect := shapes.NewShape("Rectangle")
```


