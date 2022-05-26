# About

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

Some special characters need to be escaped with a leading backslash:

| Value | Description          |
| ----- | -------------------- |
| `\a`  | Alert or bell        |
| `\b`  | Backspace            |
| `\\`  | Backslash            |
| `\t`  | Horizontal tab       |
| `\n`  | Line feed or newline |
| `\f`  | Form feed            |
| `\r`  | Carriage return      |
| `\v`  | Vertical tab         |
| `\'`  | Single quote         |
| `\"`  | Double quote         |

```go
fmt.Println("Joe\nWilliam\nJack\nAverell") 

// =>
// Joe
// William
// Jack
// Averell
```

## Raw string literals

Raw string literals use backticks (\`) as their delimiter instead of double quotes and all characters in it are interpreted literally, meaning that there is no need to escape characters or newlines.

This is an example of a multiline string:

```go
const daltons = `Joe
William
Jack
Averell`
```

This string has multiple lines. More specifically, it has 3 `\n`. However, because we are a raw string literal, we don't need to put the `\n` explicitly in the string. The newlines will be interpreted literally.

Here are some other examples comparing raw string literals with regular string literals:

```go
`abc`
// same as "abc"

"\"" // regular string literal with 1 character: a double quote
// same as `"` a raw string literal with 1 character: a double quote

`\n
` // raw string literal with 3 character: a backslash, an 'n', and a newline
// same as "\\n\n"  a regular string literal with 3 characters: a backslash, an 'n', and a newline

"\t\n" // regular string literal with 2 characters: a tab and a newline
`\t\n`// raw string literal with 4 characters: two backslashes, a 't', and an 'n'
```


