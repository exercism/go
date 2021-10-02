# Hints

## 1. Identify which application emitted a log

- The `range` keyword can be used to iterate over a given string's runes.
- Runes can be compared to other runes using an `if` conditional.
- A character within single quotes is a `rune` in Go.

## 2. Fix corrupted logs

- String concatenation can be used to build the modified log line character by character.

## 3. Determine if a log can be displayed

- Runes can be 1, 2, 3, or 4 bytes so the builtin `len` function may not accurately reflect the number of characters in a string.
