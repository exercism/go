# About

## Struct types

We've already seen struct types; to recap a `struct` is a sequence of named elements called _fields_, each field having a name and type.
The name of a field must be unique within the struct.
`Structs` can be compared with the _class_ in the Object Oriented Programming paradigm.

You create a new struct by using the `type` and `struct` keywords, then explicitly define the name and type of the fields as shown in the example below.

```go
type StructName struct{
    field1 fieldType1
    field2 fieldType2
}
```

## Non-struct types

It's also possible to define non-struct types which you can use as an alias for a built in type declarations.

```go
type Name string
func SayHello(n Name) {
  fmt.Printf("Hello %s\n", n)
}
n := Name("Fred")
SayHello(n)
// Output: Hello Fred
```

You can also define non-struct types composed of arrays and maps.

```go
type Names []string
func SayHello(n Names) {
  for _, name := range n {
    fmt.Printf("Hello %s\n", name)
  }
}
n := Names([]string{"Fred", "Bill"})
SayHello(n)
// Output:
// Hello Fred
// Hello Bill
```
