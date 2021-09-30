# Hints

## 1. Identify the application

- The `range` keyword can be used to iterate over a given string's runes.
- Runes can be compared to other runes using an `if` conditional.
- A character within single quotes is a `rune` in Go.

## 2. Redact invalid characters

- Strings are immutable in Go so you'll have to build your own string as you perform redactions.

## 3. Replace incorrect characters

- Strings are immutable in Go so you'll have to build your own string as you perform replacements.

## 4. Count characters

- Runes can be 1, 2, 3, or 4 bytes so the builtin `len` function may not accurately reflect the number of characters in a string.
