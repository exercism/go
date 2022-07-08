# Introduction

In Go, a `struct` is a sequence of named elements called _fields_, each field has a name and type.
The name of a field must be unique within the struct.
Structs can be compared with _classes_ in the Object-Oriented Programming paradigm.

## Defining a struct

You create a new struct by using the `type` and `struct` keywords, and explicitly define the name and type of the fields.
For example, here we define `Shape` struct, that holds the name of a shape and its size:

```go
type Shape struct {
    name string
    size int
}
```

Field names in structs follow the Go convention - fields whose name starts with a lower case letter are only visible to code in the same package, whereas those whose name starts with an upper case letter are visible in other packages.

## Creating instances of a struct

Once you have defined the `struct`, you need to create a new instance defining the fields using their field name
in any order:

```go
s := Shape {
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
s := Shape {
	"Oval",
	20,
}
```

However, this syntax is considered _brittle code_ since it can break when a field is added, especially when the new field is of a different type.
In the following example we add an extra field to `Shape`:

```go
type Shape struct {
	name        string
	description string // new field 'description' added
	size        int
}

s := Shape{
    "Circle",
    20,
}
// Since the second field of the struct is now a string and not an int,
// the compiler will throw an error when compiling the program:
// Output: cannot use 20 (type untyped int) as type string in field value
// Output: too few values in Shape{...}
```

To create an instance of a `struct` and get a pointer to it use a struct literal:

```go
s := &Shape {
    name: "Square",
    size: 25,
}
// s type is *Shape
```

## "New" functions

Sometimes it can be nice to have functions that help us create struct instances.
By convention, these functions are usually called `New` or have their names starting with `New`, but since they are just regular functions, you can give them any name you want.
They might remind you of constructors in other languages, but in Go they are just regular functions.

In the following example, one of these `New` functions is used to create a new instance of `Shape` and automatically set a default value for the `size` of the shape:

```go
func NewShape(name string) Shape {
	return Shape{
		name: name,
		size: 100, //default-value for size is 100
	}
}

NewShape("Triangle")
// => Shape { name: "Triangle", size: 100 }

```

Using `New` functions can have the following advantages:
* validation of the given values
* handling of default-values
* since `New` functions are often declared in the same package of the structs they initialize, they can initialize even private fields of the struct

To dive deeper into this type, you can check these resources: [Go by Example: Structs][go-by-example-structs], [A Tour of Go: Structs][tour-of-go-structs] or [Structures in Go (structs)][medium-structures-in-go]

## `new` builtin

Another way of creating a new instance of a struct is by using the `new` built-in: 

```go
s := new(Shape) // s will be of type *Shape (pointer to shape)
fmt.Printf("name: %s size: %d\n", s.name, s.size)
// Output: name:  size: 0
```

In this example, `new` creates an instance of the struct `Shape` with all the values initialized to the zero value of their type and returns a pointer to it.

~~~~exercism/note
More often than not, you will not see new instances of structs created using the `new` builtin.
Always prefer using a struct literal, a custom `New` function or any other function provided by the package that can help you build a struct.
Use the `new` builtin to create new instances of structs as a last resort.
~~~~

[go-by-example-structs]: https://gobyexample.com/structs
[medium-structures-in-go]: https://medium.com/rungo/structures-in-go-76377cc106a2
[tour-of-go-structs]: https://tour.golang.org/moretypes/2
