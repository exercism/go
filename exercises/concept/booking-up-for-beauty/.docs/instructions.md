# Instructions

In this exercise you'll be working on an appointment scheduler for a beauty salon that opened on September 15th in 2012.

You have five tasks, which will all involve appointment dates.

## 1. Parse appointment date

Implement the `Schedule` function to parse a textual representation of an appointment date into the corresponding `time.Time` format:

```go
Schedule("7/25/2019 13:45:00")
// Output: 2019-07-25 13:45:00 +0000 UTC
```

## 2. Check if an appointment has already passed

Implement the `HasPassed` function that takes an appointment date and checks if the appointment was somewhere in the past:

```go
HasPassed("July 25, 2019 13:45:00")
// Output: true
```

## 3. Check if appointment is in the afternoon

Implement the `IsAfternoonAppointment` function that takes an appointment date and checks if the appointment is in the afternoon (>= 12:00 and < 18:00):

```go
IsAfternoonAppointment("Thursday, July 25, 2019 13:45:00:00")
// Output: true
```

## 4. Describe the time and date of the appointment

Implement the `Description` function that takes an appointment date and returns a description of that date and time:

```go
Description("7/25/2019 13:45:00")
// Output: "You have an appointment on Thursday, July 25, 2019, at 13:45."
```

## 5. Return the anniversary date

Implement the `AnniversaryDate` function that returns this year's anniversary date:

```go
AnniversaryDate()
// Output: 2020-09-15
```
