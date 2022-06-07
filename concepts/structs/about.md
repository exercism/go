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
person:= Person{}
fmt.Printf("name: %s age: %d\n", person.name, person.age)
// Output: name: age: 0
```

Another way of creating a new instance of a struct is by using the `new` built-in: 
```go
s2 := new(Shape)
fmt.Printf("name: %s size: %d\n", s2.name, s2.size)
// Output: name:  size: 0
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
struct `Shape` the unexported fields `name` and `size` cannot be accessed. Field names in structs follow the Go convention - fields whose name starts
with a lower case letter are only visible to code in the same package, whereas those whose name starts with an upper case letter are visible in other packages.
Using a constructor the initial value can nevertheless be set:
```go
rect := shapes.NewShape("Rectangle")
```

You can create an instance of a `struct` without using the field names, as long as you define the fields _in order_:

```go
person := Person{
	"Mary",
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

person := Shape{
    "Circle",
    20,
}
// The compiler will fail with a similar error:
// Output: cannot use 20 (type untyped int) as type string in field value
// Output: too few values in Shape{...}
```

To dive deeper into this type, you can check these resources: [Go by example: Structs], [A Tour of Go] or [Structures in Go (structs)]

[go by example: structs]: https://gobyexample.com/structs
[structures in go (structs)]: https://medium.com/rungo/structures-in-go-76377cc106a2
[a tour of go]: https://tour.golang.org/moretypes/2
