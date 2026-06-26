# Introduction

Go is statically typed, so every variable will have a type known at compile time.
Once declared, a variable's type cannot change.

## Creating a Variable

The `var` statement can declare a new variable's type, set its value, or do both at once:

```go
var total int     // Type explicitly declared, value implicitly set (see note on zero values below)
var count int = 1 // Type explicitly declared, value explicitly set
var sum = 100     // Type implictly declared, value explicitly set
```

Inside a function, the `:=` short assignment statement can also be used.

```go
count := 1 // Type implictly declared, value explicitly set
```

A variable in Go always has a value.
If one isn't set at declaration, Go provides one [called a zero value][zero-values], based on the variable's type.

```go
var total int     // 0
var name string   // ""
var ready bool    // false
```

## Updating a Variable

Inside a function, the `=` assignment operator can set a new value for an existing variable:

```go
count := 1
count = 2
```

The new value must have the same type as the variable:

```go
count := 1
count = false // compile error: a Boolean is not an int
```

[zero-values]: https://exercism.org/tracks/go/concepts/zero-values
