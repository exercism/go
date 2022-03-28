# Hints

## 1. Calculate the value of any given card.

- The `ParseCard` function should take the `card` string (e.g. `ace`) and return its value (e.g. 11).
- Use a big [`switch` statement][switch_statement] on the `card` variable.
- King, Queen, Jack and 10 can be handled with a single case.
- The switch can have a `default` case. In any case the function should return `0` for unknown cards.

## 2. Implement the decision logic for the first turn.

- Use function `ParseCard` to determined the value for each card.
- Compute the player score by adding up the values of the two player cards
- You can either use a big [`switch` statement][switch_statement] on the player
  score (maybe with nested `if`-statements on the dealer-score in some cases),
- or you could distinguish separate player score categories (say "small hands"
  with a score less than 12, "medium hands" with a score in the range 12..20 and
  "large hands" with a score greater than 20) and write separate functions for
  all (or some) of these categories.

[logical_operators]: https://golang.org/ref/spec#Logical_operators
[if_statement]: https://golang.org/ref/spec#If_statements
[switch_statement]: https://golang.org/ref/spec#Switch_statements
