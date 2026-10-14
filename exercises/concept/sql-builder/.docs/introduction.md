# Introduction

In Go, strings are immutable.
Concatenating strings with `+` in a loop causes frequent allocations and copies data multiple times.
The [`strings` package](https://pkg.go.dev/strings) provides [`Builder`](https://pkg.go.dev/strings#Builder) to minimize memory overhead by efficiently assembling strings.

## Efficient Text Assembly

The zero-value of a `strings.Builder` is ready to use immediately without any explicit initialization or allocation:

```go
var sb strings.Builder
```

## Appending Content

You can push new data into the builder using specific write methods.
Each write method returns the number of bytes written and a nil error.
This error is typically ignored in everyday code.

The most frequently used methods are:

- `sb.WriteString(s string)`: Appends a string to the buffer
- `sb.WriteRune(r rune)`: Appends the UTF-8 encoding of a rune
- `sb.WriteByte(c byte)`: Appends a single byte

## Retrieving and Resetting

Once you have finished assembling your text, you can extract the accumulated content as a standard Go string by calling the `String()` method.
Calling `String()` is a very cheap operation.
It returns the underlying byte slice as a string directly without copying the accumulated data.

If you need to reuse the same builder instance to assemble a new string (for example, inside a loop), you can call the `Reset()` method.
This empties the internal buffer.
It allows you to start fresh and build the next string while preserving the already allocated memory capacity under the hood.

```go
sb.Reset()
```
