# Introduction

In Go, strings are immutable.
This means that every time you concatenate strings using the `+` operator, Go creates a completely new string in memory and copies the data.
While this is fine for short or infrequent operations, doing this inside loops can severely degrade performance and consume a lot of memory.

To solve this problem, Go's standard library provides the `strings.Builder` type.

## Efficient Concatenation

The `strings.Builder` type is used to efficiently build strings from parts.
It minimizes memory allocations by writing to an internal byte buffer.

To use it, you first declare a variable of type `strings.Builder`.
Its zero value is ready to use immediately:

```go
var sb strings.Builder
```

## Writing Content

You can append data to the builder using various write methods.
The most common ones are:

* `WriteString(s string)`: Appends a string.
* `WriteRune(r rune)`: Appends a single rune (Unicode code point).
* `WriteByte(c byte)`: Appends a single byte.

Each of these methods returns the number of bytes written and an error.
While this error is always `nil` (implemented to comply with the `io.Writer` interface), it is typically ignored in everyday code.

```go
sb.WriteString("Hello, ")
sb.WriteRune('🍏')
```

## Performance Tuning

If you know the size of the final string in advance, you can optimize memory usage using the `Grow(n int)` method.
This pre-allocates enough space in the internal buffer to hold `n` additional bytes.
Using `Grow` prevents the buffer from resizing during runtime, minimizing allocation overhead.

```go
var sb strings.Builder
sb.Grow(32) // Pre-allocate buffer for 32 bytes
```

## Retrieving and Resetting

Once you are done building the string, you can retrieve the accumulated text as a standard Go string using the `String()` method:

```go
result := sb.String()
```

If you need to reuse the same builder instance for a new string, you can call the `Reset()` method.
This empties the internal buffer, allowing you to start fresh without allocating a new builder.

```go
sb.Reset()
```
