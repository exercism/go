## Implementation

Define a single function, Solve, which returns two strings,
whose values are the answers to the zebra-puzzle questions
"Who drinks water?" and "Who owns the Zebra?".
Each return value will be one of the resident's nationalities:
Englishman, Spaniard, Ukrainian, Norwegian, or Japanese.

For the Solve function, use the following signature:

```
func Solve() (waterDrinker string, zebraOwner string)
```

Obviously, you could simply write a one-liner function
if you peek at the test program to see the expected return values.
But the goal is to develop an algorithm which uses
the given facts and constraints for the puzzle
and determines the two correct answers.
