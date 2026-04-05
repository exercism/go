# Introduction

Constants are variables that are set to a value at compile time and cannot be reassigned.
Use the `const` keyword to declare a constant:

```go
const greeting = "hello"
```

Multiple constants can be declared in a block.
Within a block, a constant without an explicit value repeats the previous expression:

```go
const (
    greeting    = "hello"
    buffSize    = 4096
    chunkSize   // 4096
)
```

Constants are typically declared at package level, though uncommonly they can also be declared inside a function.

## Typed and untyped constants

Declaring a constant without specifying a type makes it untyped.
An untyped constant has no set type, but its value determines what types it's compatible with.
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

Constants can only hold booleans, numbers, strings, and runes.
Because the values are set at compile time, the results of a function call can not be constants:

```go
const s = math.Sqrt(4)  // compile error because the value isn't known at compile time
```

## Iota

Within a block of constants, `iota` represents each constant's position within the block:

```go
const (
    Sunday    = iota // 0
    Monday           // 1
    Tuesday          // 2
    Wednesday        // 3
)
```

Each constant after the first inherits the same expression, with `iota` incrementing by one.
