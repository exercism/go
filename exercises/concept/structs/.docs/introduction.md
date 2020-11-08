## structs

In Go, a `struct` is a sequence of named elements called _fields_, each field having a name and type. The name of a field must be unique within the struct. `Structs` can be compared with the _class_ in the Object Oriented Programming paradigm.

You create a new struct by using the `struct` keyword, a **_built-in type_**, and explicty define the name and type of the fields as shown in the example below.

```go
type StructName struct{
    field1 fieldType1
    field2 fieldType2
}
```

Struct fields are accessed using a `.` notation.

```go
someStruct.someField = "someValue"
```
