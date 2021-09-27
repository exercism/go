# Instructions

You have been tasked with parsing logs from multiple applications so they can be identified, cleaned, and stored in a log aggregation system.

## 1. Identify the application

Logs come from multiple applications that each use their own proprietary log format. The application emitting a log must be identified before it can be stored in the log aggregation system.

Implement the `LogApplication` function that takes one argument:

- A log line

The function should identify the application that is emitting the log line and return it to the caller.

An application can be identified by searching for a specific sequence of characters in the log line. The following table details which specific sequence of characters corresponds to a given application.

| Application      | Character Sequence | Unicode Code Points  |
|------------------|--------------------|----------------------|
| `recommendation` | ❗❗               | `U+2757`, `U+2757`   |
| `search`         | 🔍🔎               | `U+1F50D`, `U+1F50E` |
| `weather`        | ☁☀                 | `U+2601`, `U+2600`   |

When a log line does not contain one of the character sequences above, return `default` to the caller. If a log line contains more than one character sequence, return the application corresponding to the first character sequence in the log line starting from left to right.

```go
LogApplication("forecast: ☁☀ rain")
// Output: weather
```

```go
LogApplication("❗❗ recommended search product 🔍🔎")
// Output: recommendation
```

```go
LogApplication("search complete")
// Output: default
```

## 2. Clean the log line

Logs may contain invalid, illegal, or even profane characters. As such, it may be necessary to clean the logs before they are stored in a log aggregation system.

Implement the `LogClean` function that takes two arguments:

- A log line
- A slice of invalid characters

The function should remove any invalid characters and return the cleaned log line to the caller.

```go
logLine := "please clean '👎' characters: 👎 thank you"
invalidChars := []rune{'👎'}

LogClean(logLine, invalidChars)
// Output: please clean '' characters:  thank you
```

```go
logLine := "XXplease clean '👎' and 'X' characters: 👎 thank youXX"
invalidChars := []rune{'👎', 'X'}

LogClean(logLine, invalidChars)
// Output: please clean '' and '' characters:  thank you"
```
