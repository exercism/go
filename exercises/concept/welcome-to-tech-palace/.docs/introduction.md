# Introduction

A `string` in Go is an immutable sequence of bytes, which don't necessarily have to represent characters.

A `string` literal is a sequence of bytes enclosed in double quotes:

```go
firstName := "Jane"
```

Strings can be concatenated via the `+` operator:

```go
fullName := "Jane" + " " + "Austen"
```

Some special characters need to be escaped with a leading backspace:

```go
question := "\t or spaces?"
```

The `strings` package contains many useful functions to work on strings.

```go
import "strings"

strings.ToLower("TEST") // Output: "test"
```
