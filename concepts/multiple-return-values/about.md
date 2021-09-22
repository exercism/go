# About Multiple Return Values

Go functions and methods can return multiple values.
Very often, a second return value is used to return an error.
For example:

```go
func GetFoo() (Foo, error) { ... }
```

The assignment for multiple return values just uses a comma to separate the variables:

```go
foo, err := GetFoo()
```

If statements can use an initializer before the condition separated by a semicolon.
This is a common idiom seen for error handling:

```go
if foo, err := GetFoo(); err != nil {
    // handle the error
}
```
