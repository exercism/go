# About Type Assertions

Interfaces in Go can introduce ambiguity about the underlying type.
In fact, the empty interface can take any concrete value at all (including primitives).
A type assertion allows us to extract the interface value's underlying concrete value using this syntax: `interfaceVariable.(concreteType)`.

If we assign the result of this statement to a single value, we assert that the interface variable holds a value of the concrete type.
As a result, it may panic at runtime if the interface value does not have the specified concrete type.
For example:

```go
var input interface{} = 12
str := input.(string) // panic at runtime since input is not a string!
```

We can test whether an interface value holds a specific concrete type by making use of both return values of the type assertion: the underlying value and a boolean value that reports whether the assertion succeeded.
For example:

```go
str, ok := input.(string) // no panic if input is not a string
```

If `input` holds a `string`, then `str` will be the underlying value and `ok` will be true.
If `input` does not hold a `string`, then `str` will be the zero value of type `string` (ie. `""` - the empty string) and `ok` will be false.
No panic occurs in any case.

It is common to see this sort of idiom:

```go
str, ok := input.(string)
if !ok {
    str = "a default value"
}
```

## Type Switches

A **type switch** can perform several type assertions in a row.
It has the same syntax as a type assertion (`interfaceVariable.(concreteType)`), but instead of a specific `concreteType` it uses the keyword `type`.
Here is an example:

```go
var i interface{} = 12 // try: 12.3, true, int64(12), []int{}, map[string]int{}

switch v := i.(type) {
case int:
    fmt.Printf("the integer %d\n", v)
case string:
    fmt.Printf("the string %s\n", v)
default:
    fmt.Printf("type, %T, not handled explicitly: %#v\n", v, v)
}
```
