# Hints

## 1. Calculate Duration

Think about how to measure the time between two moments.  
Go provides a built-in way to compare timestamps directly.

---

## 2. Handle Different Time Zones

You don’t need to manually adjust for time zones.  
Focus on comparing the two timestamps as they are.

---

## 3. Parse Recorded Times

The input follows a structured format with hours, minutes, and seconds.  
Look for a standard library function that can interpret duration strings.

---

## 4. Correct Start Time

You need to adjust a timestamp by moving it backward in time.  
Consider how to shift a time value using a duration.

---

## 5. Filter Valid Detections

First determine how much time has passed since the start.  
Then check whether that elapsed time falls within a given range.

---

## 6. Format Leaderboard Output

You need to display:
- the runner’s total time
- the difference compared to the leader

Think about:
- how to convert a duration to a readable string
- how to express the difference in seconds with precision

---

## General Advice

- Prefer built-in time utilities instead of manual calculations
- Be careful with boundaries when checking ranges
- Keep your solution simple and precise — timing systems must be reliable