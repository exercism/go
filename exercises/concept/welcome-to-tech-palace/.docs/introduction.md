# Introduction

A `string` in Go is an immutable sequence of bytes, which don't necessarily have to represent characters.

A string literal is defined between double quotes:

```go
const name = "Jane"
```

Strings can be concatenated via the `+` operator:

```go
fullName := "Jane" + " " + "Austen"
// "Jane Austen"
```

Some special characters need to be escaped with a leading backslash, such as `\t` for a tab and `\n` for a new line in strings.

```go
Dialogue := "How is the weather today?\nIt's sunny"  
// "How is the weather today?
// It's sunny"
```

The `strings` package contains many useful functions to work on strings.
For more information about string functions, check out the [strings package documentation](https://pkg.go.dev/strings).
Here are some examples:

```go
import "strings"

// strings.ToLower returns the string given as argument with all its characters lowercased
strings.ToLower("MaKEmeLoweRCase")
// "makemelowercase"

// strings.Repeat returns a string with a substring given as argument repeated many times
strings.Repeat("Go", 3)
// "GoGoGo"
```
