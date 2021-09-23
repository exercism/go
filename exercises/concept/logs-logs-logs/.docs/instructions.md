# Instructions

In this exercise you'll be processing log-lines.

Each log line is a string formatted as follows: `"[<LEVEL>]: <MESSAGE>"`.

There are three different log levels:

- `INFO`
- `WARNING`
- `ERROR`

You have three tasks, each of which will take a log line and ask you to do something with it.

## 1. Get message from a log line

Implement the `Message` function to return a log line's message:

```go
Message("[ERROR]: Invalid operation")
// Output: "Invalid operation"
```

Any leading or trailing white space should be removed:

```go
Message("[WARNING]:  Disk almost full\r\n")
// Output: "Disk almost full"
```

## 2. Get the message length in characters

Implement the `MessageLen` function to return a log line's message length:

```go
MessageLen("[ERROR]: Invalid operation \n")
// Output: 17
```

## 3. Get log level from a log line

Implement the `LogLevel` function to return a log line's log level, which should be returned in lowercase:

```go
LogLevel("[ERROR]: Invalid operation")
// Output: "error"
```

## 4. Reformat a log line

Implement the `Reformat` function that reformats the log line, putting the message first and the log level after it in parentheses:

```go
Reformat("[INFO]: Operation completed")
// Output: "Operation completed (info)"
```
