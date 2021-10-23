# Introduction

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
