# Introduction

In Go, strings are immutable. This means that every time you concatenate strings using the `+` operator, Go creates a completely new string in memory and copies the data. While this is fine for short or infrequent operations, doing this inside loops can severely degrade performance and consume a lot of memory.

To solve this problem, Go's standard library provides the `strings.Builder` type.

## Efficient Concatenation

The `strings.Builder` type is used to efficiently build strings from parts. It minimizes memory allocations by writing to an internal byte buffer.

To use it, you first declare a variable of type `strings.Builder`. Its zero value is ready to use immediately:

```go
var builder strings.Builder
```

## Writing Content

You can append data to the builder using various write methods. The most common ones are:

* `WriteString(s string)`: Appends a string.
* `WriteRune(r rune)`: Appends a single rune (Unicode code point).
* `WriteByte(c byte)`: Appends a single byte.

Each of these methods returns the number of bytes written and an error (which is part of the `io.Writer` interface compliance). However, `strings.Builder` is guaranteed to never return an error (it is always `nil`), so you can safely ignore these return values in everyday code.

```go
builder.WriteString("Hello, Go ")
builder.WriteRune('🌎')
```

## Retrieving and Resetting

Once you are done building the string, you can retrieve the accumulated text as a standard Go string using the `String()` method:

```go
result := builder.String() // Returns "Hello, Go 🌎"
```

If you need to reuse the same builder instance for a new string, you can call the `Reset()` method. This empties the internal buffer, allowing you to start fresh without allocating a new builder structure in memory.

```go
builder.Reset()
```
