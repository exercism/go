# About

In Go, a `struct` is a sequence of named elements called _fields_, each field having a name and type. The name of a field must be unique within the struct. `Structs` can be compared with the _class_ in the Object Oriented Programming paradigm.

You create a new struct by using the `struct` keyword and explicty define the name and type of the fields as shown in the example below. An `empty struct` is a struct with all fields set to their own zero values.

```go
type StructName struct{
    field1 fieldType1
    field2 fieldType2
}
```

To access the value of a field from a struct, we us the `.` operator. This way you can set or get values of a struct field.

```go
type Animal struct{
    name string
    age int
}

// create a new variable
var dog Animal

// assign a values to each field by using the . operator
dog.name = "Rex"
dog.age = 10

// or you can create a struct with field values
dog := Animal{
    name: "Bronson",
    age: 7,
}
```

Fields that don't have a initial value assigned, will have their zero value.

To dive deeper into this type, you can check these resources: [Go by example: Structs], [A Tour of Go] or [Structures in Go (structs)]

[go by example: structs]: https://gobyexample.com/structs
[structures in go (structs)]: https://medium.com/rungo/structures-in-go-76377cc106a2
[a tour of go]: https://tour.golang.org/moretypes/2
