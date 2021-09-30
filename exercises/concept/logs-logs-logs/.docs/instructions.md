# Instructions

You have been tasked with creating a log library to assist with managing your organization's logs. This library will provide functionality to identify which application emitted a given log, to redact and replace information from logs, and to count the number of characters within a log line.

## 1. Identify the application

Logs come from multiple applications that each use their own proprietary log format. The application emitting a log must be identified before it can be stored in the log aggregation system.

Implement the `Application` function using the following function signature:

```
func Application(log string) string
```

The `Application` function should read the string `log` and identify which application emitted the log.

To identify which application emitted a given log line, search the log line for a specific character as specified by the following table:

| Application      | Character | Unicode Code Point |
|------------------|-----------|--------------------|
| `recommendation` | ❗        | `U+2757`           |
| `search`         | 🔍        | `U+1F50D`          |
| `weather`        | ☀         | `U+2600`           |

If a log line does not contain one of the characters from the above table, return `default` to the caller. If a log line contains more than one character in the above table, return the application corresponding to the first character found in the log line starting from left to right.

```go
Application("forecast: ☀ sunny")
// Output: weather
```

```go
Application("❗ recommended search product 🔍")
// Output: recommendation
```

```go
Application("search complete")
// Output: default
```

## 2. Redact invalid characters

Logs may contain invalid, illegal, or even profane characters. As such, it may be necessary to redact information from the logs before they are stored in a log aggregation system.

Implement the `Redact` function using the following function signature:

```
func Redact(log string, redactions []rune) string
```

The `Redact` function should read the `log`, redact all occurances of the characters in `redactions` from `log`, and return the redacted log line to the caller.

```go
log := "please redact '👎' characters: 👎 thank you"
redactions := []rune{'👎'}

Redact(log, redactions)
// Output: please redact '' characters:  thank you
```

```go
log := "XXplease redact '👎' and 'X' characters: 👎 thank youXX"
redactions := []rune{'👎', 'X'}

Redact(log, redactions)
// Output: please redact '' and '' characters:  thank you"
```

## 3. Replace incorrect characters

Logs may contain incorrect characters. As such, these characters may need to be replaced with other characters before the log can be stored in a log aggregation system.

Implement the `Replace` function using the following function signature:

```
func Replace(log string, old, new rune) string
```

The `Replace` function should read the `log`, replace all occurances of the `old` character with the `new` character within `log`, and return the modified log line to the caller.

```go
log := "please replace '👎' with '👍'"

Replace(log, '👎', '👍')
// Output: please replace '👍' with '👍'"
```

```go
log := "please replace 'e' with 'X'"

Replace(log, 'e', 'X')
// Output: plXasX rXplacX 'X' with 'X'"
```

## 4. Count characters

After using this library, users are asking for the ability to count the number of characters in log lines.

Implement the `Count` function using the following function signature:

```
func Count(log string) int
```

The `Count` function should read the `log` and return the number of characters in the log line.

```go
Count("exercism is fun")
// Output: 15
```

```go
Count("🧠exercism is fun❗")
// Output: 17
```
