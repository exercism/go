# Introduction

In Go, a constant is a simple, unchanging value assigned to a name with the `const` keyword:

```go
const myWebsite = "Exercism"
```

Such constants are untyped, but they are given a default type based on their syntax when a type is required, such as when they are passed to a method. Typed constants can be created by explicitly adding a type:

```go
const myWebsite string = "Exercism"
```

Go does not have enums like some other languages, but does have a predeclared identifier called `iota` for creating enumerated constants. Constants in a block are implicitly repeated:

```go
const (
    a = 9
    b
    c
    d = iota
    e
    f
    g
)
fmt.Print(a, b, c, d, e, f, g)
// Output: 9 9 9 3 4 5 6
```
