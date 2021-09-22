# About

A `string` in Go is an immutable sequence of bytes, which don't necessarily have to represent characters.

A string literal is defined between double quotes:

```go
const name = "Jane"
```

Some values need to be escaped:

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
const daltons = "Joe\nWilliam\n\nJack\nAverell`"
```

Raw string literals use backticks (`) as their delimiter instead of double quotes and are interpreted literally, meaning that there is no need to escape characters or newlines.

```go
const daltons = `Joe
William
Jack
Averell`
```
