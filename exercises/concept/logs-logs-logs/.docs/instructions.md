# Instructions

You have been tasked with creating a log library to assist with managing your organization's logs. This library will allow users to identify which application emitted a given log, to fix corrupted logs, and to determine if a given log line is within a certain character limit.

## 1. Identify which application emitted a log

Logs come from multiple applications that each use their own proprietary log format. The application emitting a log must be identified before it can be stored in a log aggregation system.

Implement the `Application` function that takes a log line and returns the application that emitted the log line.

To identify which application emitted a given log line, search the log line for a specific character as specified by the following table:

| Application      | Character | Unicode Code Point |
|------------------|-----------|--------------------|
| `recommendation` | ❗        | `U+2757`           |
| `search`         | 🔍        | `U+1F50D`          |
| `weather`        | ☀         | `U+2600`           |

If a log line does not contain one of the characters from the above table, return `default` to the caller. If a log line contains more than one character in the above table, return the application corresponding to the first character found in the log line starting from left to right.

```go
Application("❗ recommended search product 🔍")
// Output: recommendation
```

## 2. Fix corrupted logs

Due to a rare but persistent bug in the logging infrastructure, certain characters in logs can become corrupted. After spending time identifying the corrupted characters and their original value, you decide to update the log library to assist in fixing corrupted logs.

Implement the `Replace` function that takes a log line, a corrupted character, and the original value and returns a modified log line that has all occurances of the corrupted character replaced with the original value.

```go
log := "please replace '👎' with '👍'"

Replace(log, '👎', '👍')
// Output: please replace '👍' with '👍'"
```

## 3. Determine if a log can be displayed

Systems responsible for displaying logs have a limit on the number of characters that can be displayed per log line. As such, users are asking for this library to include a helper function to determine whether or not a log line is within a specific character limit.

Implement the `WithinLimit` function that takes a log line and character limit and returns whether or not the log line is within the character limit.

```go
WithinLimit("hello❗", 6)
// Output: true
```
