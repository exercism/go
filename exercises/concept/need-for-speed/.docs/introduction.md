# Introduction

In Go, a `struct` is a sequence of named elements called _fields_, each field has a name and type. The name of a field must be unique within the struct.
`Structs` can be compared with _classes_ in the Object-Oriented Programming paradigm.

You create a new struct by using the **_built-in type_** keywords `type` and `struct`, and explicitly define the name and type of the fields.
For example, to define a `Person` struct:

```go
type Person struct{
    name string
    age int
}
```

Once you have defined the `struct`, you need to create a new instance defining the fields _in order_, for example:

```go
person := Person{
	"John",
	20,
}
```

Or you can explicitly define the field name, in any order:

```go
anotherPerson := Person{
	age: 22,
	name: "Anna",
}
```

To access instance fields, use the `.` notation. For example, to print the name and age of a `Person`
you can do the following:

```go
fmt.Printf("name: %s age: %d\n", person.name, person.age)
// Output: name: John age: 22
```

Instance fields can be modified by accessing each field separately:

```go
anotherPerson.name = "Mary"
```
