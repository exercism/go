# Introduction

In Go, strings are immutable. This means that once a string is created, its content cannot be changed. Any standard modification or concatenation (such as using the `+` operator in a loop) forces Go to allocate a completely new string in memory and copy all the data over from the old one. For frequent operations, this behavior drastically reduces performance and causes excessive memory consumption.

To solve this efficiency problem, Go's standard library provides the `strings.Builder` type in the `strings` package.

## Efficient Text Assembly

The `strings.Builder` type accumulates text inside an internal, mutable byte buffer, minimizing memory allocation and copying overhead. 

Its zero-value is immediately ready to use without any explicit initialization or allocation:

```go
var builder strings.Builder
```

## Appending Content

You can push new data into the builder using specific write methods. The most frequently used methods are:

- `builder.WriteString(s string)`: Appends the contents of a string.
- `builder.WriteRune(r rune)`: Appends a single Unicode character.
- `builder.WriteByte(c byte)`: Appends a single raw byte.

## Retrieving Results

Once you have finished assembling your text, you can extract the accumulated content as a standard Go string by calling the `String()` method:

```go
result := builder.String()
```
