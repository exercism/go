# Instructions

This exercise addresses the parsing of log files.

After a recent security review you have been asked to clean up the organization's archived log files.

All strings passed to the functions are guaranteed to be non-null and without leading and trailing spaces.

## 1. Identify garbled log lines

You need some idea of how many log lines in your archive do not comply with current standards.
You believe that a simple test reveals whether a log line is valid.
To be considered valid a line should begin with one of the following strings:

- [TRC]
- [DBG]
- [INF]
- [WRN]
- [ERR]
- [FTL]

Implement the `IsValidLine` function to return `false` if a string is not valid otherwise `true`.

```go 
IsValidLine("[ERR] A good error here");
// => true
IsValidLine("Any old [ERR] text");
// => false
IsValidLine("[BOB] Any old text");
// => false
```

## 2. Split the log line

A new team has joined the organization, and you find their log files are using a strange separator for "fields".
Instead of something sensible like a colon ":" they use a string such as "<--->" or "<=>" (because it's prettier) in fact any string that has a first character of "<" and a last character of ">" and any combination of the following characters "~", "\*", "=" and "-" in between.

Implement the `SplitLogLine` function that takes a line and returns an array of strings each of which contains a field.

```go
SplitLogLine("section 1<*>section 2<~~~>section 3")
// => []string{"section 1", "section 2", "section 3"},
```

## 3. Count the number of lines containing `password` in quoted text

The team needs to know about references to passwords in quoted text so that they can be examined manually.

Implement the `CountQuotedPasswords` function to provide an indication of the likely scale of the manual exercise.

Identify log lines where the string "password", which may be in any combination of upper or lower case, is surrounded by quotation marks.
You should account for the possibility of additional content between the quotation marks before and after "password".
Each line will contain at most two quotation marks.

Lines passed to the routine may or may not be valid as defined in task 1.
We process them in the same way, whether or not they are valid.

```go
lines := []string{
    `[INF] passWord`, // contains 'password' but not surrounded by quotation marks
    `"passWord"`,  // count this one
    `[INF] User saw error message "Unexpected Error" on page load.`, // does not contain 'password'
    `[INF] The message "Please reset your password" was ignored by the user`, // count this one 
}
// => 2
```

## 4. Remove artifacts from log

You have found that some upstream processing of the logs has been scattering the text "end-of-line" followed by a line number (without an intervening space) throughout the logs.

Implement the `RemoveEndOfLineText` function to take a string and remove the end-of-line text and return a "clean" string.

Lines not containing end-of-line text should be returned unmodified.

Just remove the end of line string. Do not attempt to adjust the whitespaces.

```go
RemoveEndOfLineText("[INF] end-of-line23033 Network Failure end-of-line27")
// => "[INF]  Network Failure "
```

## 5. Tag lines with user names

You have noticed that some of the log lines include sentences that refer to users.
These sentences always contain the string `"User"`, followed by one or more space characters, and then a user name.
You decide to tag such lines.

Implement a function `TagWithUserName` that processes log lines:

- Lines that do not contain the string `"User "` remain unchanged.
- For lines that contain the string `"User "`, prefix the line with `[USR]` followed by the user name.
 
For example:

```go
result := TagWithUserName([]string{
    "[WRN] User James123 has exceeded storage space.",
	"[WRN] Host down. User   Michelle4 lost connection.",
	"[INF] Users can login again after 23:00.",
	"[DBG] We need to check that user names are at least 6 chars long.",
     
}) 
// => []string {
//  "[USR] James123 [WRN] User James123 has exceeded storage space.",
//  "[USR] Michelle4 [WRN] Host down. User   Michelle4 lost connection.",
//  "[INF] Users can login again after 23:00.",
//  "[DBG] We need to check that user names are at least 6 chars long."
// }
```

You can assume that: 

- User names are followed by at least one whitespace character in the log.
- There is at most one occurrence of the string `"User "` in each line.
- User names are non-empty strings that do not contain whitespace.

 
