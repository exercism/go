# About
In Go, a `struct` is a sequence of named elements called _fields_, each field has a name and type. The name of a field must be unique within the struct.
Structs can be compared with _classes_ in the Object-Oriented Programming paradigm.

You create a new struct by using the `type` keyword and the **_built-in type_** `struct`, and explicitly define the name and type of the fields.
For example, to define a `Person` struct:

```go
type Person struct{
    name string
    age int
}
```

Once you have defined the `struct`, you need to create a new instance defining the fields using their field name
in any order:

```go
person := Person{
	name: "John",
	age: 22,
}
```

To read or modify instance fields, use the `.` notation:

```go
// Update the age
person.age = 23
fmt.Printf("name: %s age: %d\n", person.name, person.age)
// Output: name: John age: 23
```

Fields that don't have an initial value assigned, will have their zero value. For example:

```go
person:= Person{}
fmt.Printf("name: %s age: %d\n", person.name, person.age)
// Output: name: age: 0
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
type Person struct {
	name string
	email string
	age int
}

person := Person{
    "Mary",
    20,
}
// The compiler will fail with a similar error:
// Output: cannot use 20 (type untyped int) as type string in field value
// Output: too few values in Person{...}
```

To dive deeper into this type, you can check these resources: [Go by example: Structs], [A Tour of Go] or [Structures in Go (structs)]

[go by example: structs]: https://gobyexample.com/structs
[structures in go (structs)]: https://medium.com/rungo/structures-in-go-76377cc106a2
[a tour of go]: https://tour.golang.org/moretypes/2
