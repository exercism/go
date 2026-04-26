# Introduction

Constants are values set at compile time and cannot be reassigned.
Use the `const` keyword to declare a single constant:

```go
const greeting = "hello"
```

Multiple constants can be declared in a block.
Within a block, a constant without an explicit value set repeats the previous expression:

```go
const (
    greeting    = "hello"
    buffSize    = 4096
    chunkSize   // 4096
)
```

Constants are typically declared at package level, though they can also be declared inside a function.

## Typed and untyped constants

Declaring a constant without specifying a type makes it untyped.
An untyped constant has no set type; its value determines what types it's compatible with.
As long as the underlying value is compatible, the compiler will convert it to the required type at each use:

```go
const numA = 2

math.Sqrt(numA)  // numA is treated as float64 here
var numB int = numA  // numA is treated as int here
```

A typed constant fixes the type at declaration:

```go
const num int = 2

math.Sqrt(num)  // compile error because a float is expected
```

## What can be a constant

Constants can only hold Booleans, numbers, strings, and runes.
Because the values are set at compile time, the results of a function call can not be used for a constant.

```go
const s = math.Sqrt(4)  // compile error because the value isn't known at compile time
```

## Iota

Within a block of constants, `iota` represents a constant's position within the block:

```go
const (
    Sunday    = iota // 0
    Monday           // 1
    Tuesday          // 2
    Wednesday        // 3
)
```

Each constant after the first inherits the same expression, with `iota` incrementing by one.

`iota` also works in expressions.

```go
const (
    _  = iota             // position 0: value discarded
    KB = 1 << (10 * iota) // position 1: 1 << 10 = 1024
    MB                    // position 2: 1 << 20 = 1048576
    GB                    // position 3: 1 << 30 = 1073741824
)
```

For a deeper look at constants in Go, see [The Go Blog][const-blog] and [Effective Go][effective-go].

[const-blog]: https://go.dev/blog/constants
[effective-go]: https://go.dev/doc/effective_go#constants
