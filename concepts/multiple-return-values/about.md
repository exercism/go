# About Multiple Return Values

Go functions and methods can return multiple values.
Very often, a second return value is used to return an error.
For example:

```go
func GetCard() (Card, error) { ... }
```

The assignment for multiple return values just uses a comma to separate the variables:

```go
card, err := GetCard()
```

If statements can use an initializer before the condition separated by a semicolon.
This is a common idiom seen for error handling:

```go
if card, err := GetCard(); err != nil {
    // handle the error
}
```
