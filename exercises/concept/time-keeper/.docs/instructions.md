# Instructions

Sarah is a professional time-keeper who travels around the world to measure athletes' performance at sporting events.

In this exercise, you will implement several helper functions for her timing software.  
Accuracy is critical — even small mistakes could invalidate a world record.
---

## 1. Calculate the race duration

Given a start time and a finish time, calculate how long the race lasted.

Implement the function `GetDuration(start, finish)` that returns the duration between the two timestamps.

```go
duration := GetDuration(start, finish)
// => 1h5m19.234s
```

## 2. Handle timestamps from different time zones

Timing devices may be located in different time zones.
You still need to calculate the correct duration between two timestamps.

Implement the function `GetDurationInDifferentTimezones(start, finish)` that returns the correct duration regardless of time zone differences.
```go
duration := GetDurationInDifferentTimezones(start, finish)
// => correct duration
```

## 3. Parse recorded race times
Some race times are recorded manually in a string format such as:
```go
"10h5m19.234s"
```

Implement the function `ParseRun(input)` that converts this string into a duration.
```go
duration, _ := ParseRun("10h5m19.234s")
// => 10h5m19.234s
```

### 4. Correct the start time
A backup system shows that the race actually started earlier than recorded.
You need to adjust the start time accordingly.

Implement the function `FixStartTime(incorrectStart, offset)` that returns the corrected start time by subtracting the given duration.
```go
newStart := FixStartTime(incorrectStart, offset)
// => corrected start time
```

## 5. Filter valid detection times

During the race, many timestamps are recorded, including unwanted ones.
Only detections within a specific time window after the start should be considered valid.

Implement the function `IsValidDetection(start, from, to, detection)` that returns whether a detection is within the allowed range.
```go
valid := IsValidDetection(start, 10*time.Minute, 30*time.Minute, detection)
// => true or false
```
## 6. Format leaderboard results
Race results should be displayed in a clear format, including the time difference from the leader.

Implement the function FormatResult(result, leader) that returns a string formatted like:
```go
"1h5m19.234s +4.843"
```

Where:
- the first part is the athlete’s time
- the second part is the difference (in seconds) compared to the leader
- 
```go
output := FormatResult(result, leader)
// => "1h5m19.234s +4.843"
```

## Notes
- Use Go’s time package for all time-related operations
- Focus on correctness — small mistakes can lead to invalid race results
- Keep your implementation simple and readable

