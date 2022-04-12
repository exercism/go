# About

A `string` in Go is an immutable sequence of bytes, which don't necessarily have to represent characters.

A string literal is defined between double quotes:


```go
const name = "Jane"
```
Strings can be concatenated via the `+` operator:

```go

fullName := "Jane" + " " + "Austen" // "Jane Austen"
```
The `strings` package contains many useful functions to work on strings, such as `Repeat`, `ToUpper`, `ToLower`, `ReplaceAll`, `TrimSpace`, and many more. For getting more information about strings functions check out the [strings package documentation](https://pkg.go.dev/strings).

```go
import "strings"

// strings.ToLower returns the string given as argument with all its characters lowercased
fmt.Println(strings.ToLower("MaKEmeLoweRCase"))
// Output: "makemelowercase"

// strings.ToUpper(s string) returns the string given as argument with all its characters uppercased
fmt.Println(strings.ToUpper("MaKemeUpPercaSe")) // Output: "MAKEMEUPPERCASE"

// Returns a new string consisting of count copies of the string s by (strings.Repeat(s string, n int))
fmt.Println(strings.Repeat("Go", 3)) // Output: "GoGoGo" 

// Returns a copy of the string s with all non-overlapping instances of old replaced by new by (strings.ReplaceAll(s, old, new string))
fmt.Println(strings.ReplaceAll("your cat is playing with your pillow", "your", "my")) // Output: "my cat is playing with my pillow

// Returns a slice of the string s, with all leading and trailing white space removed by (strings.TrimSpace(s string))
fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n")) // Output: "Hello, Gophers"

```

Some special characters need to be escaped with a leading backslash;

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

// Output:
// Joe
// William
// Jack
// Averell
```

Raw string literals use backticks (\`) as their delimiter instead of double quotes and are interpreted literally, meaning that there is no need to escape characters or newlines.

```go
const daltons = `Joe
William
Jack
Averell`
```


