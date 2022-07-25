# Introduction

In Go, functionality for working with times is provided by the `time` package. The types and methods in this package allow us to manipulate times, get the current time, determine elapsed time, parse times from strings, and more.

To work with time, you will usually call a method on a `Time` instance, but there are also some functions called on the `time` package itself.

Date-time strings are represented a bit differently in Go. In many languages and libraries, developers represent their desired date-time string format with the [ISO-8601 standard](https://en.wikipedia.org/wiki/ISO_8601). In this syntax, developers must express their input/output date-time string like `YYYY mm:ss`. If developers want the library to represent `5:43 PM` as `17:43` they pass the format string `"HH:mm"`. The `time` package in Go, allows developers to instead represent a specific time in their desired format.

 Instead of requiring developers to remember what capital vs lower case letters represent, the `time` package allows users to represent a specific date-time in history in exactly their desired format. Developers manually format the go-specified date-time, `Mon Jan 2 15:04:05 -0700 MST 2006`, to their desired format to parse or format strings. If they want `5:43 PM` formatted as `17:43` they pass `"15:04"` as the date-time format string since they have to represent 3pm and 4 minutes from the specified reference date. The reason this specific date-time was chosen is that it can be represented as
 
 - `1` (Jan)
 - `2` (the 2nd)
 - `3` (3PM)
 - `4` (04 minute)
 - `5` (5th second)
 - `6` (year 2006)
 - `7` (-7 hours from GMT)
