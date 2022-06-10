# About

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

Field names in structs follow the Go convention - fields whose name starts with a lower case letter are only visible to code in the same package, whereas those whose name starts with an upper case letter are visible in other packages.

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

Fields that don't have an initial value assigned, will have their zero value. For example:

```go
s := Shape{name: "Circle"}
fmt.Printf("name: %s size: %d\n", s.name, s.size)
// Output: name: Circle size: 0
```

You can create an instance of a `struct` without using the field names, as long as you define the fields _in order_:

```go
s := Shape{
	"Oval",
	20,
}
```

However, this syntax is considered _brittle code_ (breaks when a field is added), especially when the new field
is of a different type:

```go
type Shape struct {
	name        string
	description string
	size        int
}

s := Shape{
    "Circle",
    20,
}
// The compiler will fail with a similar error:
// Output: cannot use 20 (type untyped int) as type string in field value
// Output: too few values in Shape{...}
```

Another way of creating a new instance of a struct is by using the `new` built-in: 

```go
s := new(Shape)
fmt.Printf("name: %s size: %d\n", s2.name, s2.size)
// Output: name:  size: 0
```

In this example, `new` creates an instance of the struct `Shape` with all the values initialized to the zero value of their type and returns a pointer to it.


Sometimes it can be nice to have functions that help us create struct instances. By convention, these functions are usually called `New` or have their names starting with `New`, but since they are just regular functions, you can give them any name you want. They might remind you of constructors in other languages, but in Go they are just regular functions.
In the following example, one of these `New` functions is used to create a new instance of `Shape` and automatically set a default value for the `size` of the shape:

```go
func NewShape(name string) Shape {
	return Shape{
		name: name,
		size: 100, //default-value for size is 100
	}
}
```

Using `New` functions can have the following advantages:
* validation of the given values
* handling of default-values
* since `New` functions are often declared in the same package of the structs they initialize, they can initialize even private fields of the struct

To dive deeper into this type, you can check these resources: [Go by example: Structs], [A Tour of Go] or [Structures in Go (structs)]

[go by example: structs]: https://gobyexample.com/structs
[structures in go (structs)]: https://medium.com/rungo/structures-in-go-76377cc106a2
[a tour of go]: https://tour.golang.org/moretypes/2
