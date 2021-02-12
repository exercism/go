# Hints

## General

Conditionals are used to check for certain conditions and/or criteria. The most basic way of performing a conditional operation is using a single `if` statement.

## 1. Calculate the score of any given card.

The `ParseCard` function should take the `card` string (e.g. `ace`) and turn it into its value (e.g. 11).

- Use a big [`switch` statement][switch_statement] on the `card` variable.
- King, Queen, Jack and 10 can be handled with a single case.
- The switch can have a `default` case. In any case the function should return `0` for unknown cards.

## 2. Determine if two cards make up a Blackjack.

`IsBlackJack` checks if 2 cards have the combined value of 21.

- Should use the `ParseCard` function to get the value for each card.
- Should sum up the values of the 2 cards.
- Should return `true` if the sum is equal to `21`.
- No `if` statement is needed here. The result for the comparison can be returned.

## 3. Implement the decision logic for hand scores larger than 20 points.

As the `LargeHand` function is only called for hands with a value larger than 20, there are only 2 different possible hands: A **BlackJack** with a total value of `21` and **2 Aces** with a total value of `22`.

- The function should check [if][if_statement] `isBlackJack` is `true` and return "P" otherwise.
- If `isBlackJack` is `true`, the dealerScore needs to be checked for being lower than 10. [If][if_statement] it is lower, return "W" otherwise "S".

## 4. Implement the decision logic for hand scores with less than 21 points.

The `SmallHand` function is only called if there are no Aces on the hand (`handScore` is less than 21).

- Implement every condition using [logical operators][logical_operators] if necessary:
  - [If][if_statement] your cards sum up to 17 or higher you should always _stand_.
  - [If][if_statement] your cards sum up to 11 or lower you should always _hit_.
  - [If][if_statement] your cards sum up to a value within the range [12, 16] you should always _stand_ if the dealer has a 6 or lower.
  - [If][if_statement] your cards sum up to a value within the range [12, 16] you should always _hit_ if the dealer has a 7 or higher.
- (optional) Try to optimize the conditions:
  - Pull together the conditions for _stand_ into one.
  - Pull together the conditions for _hit_ into one.
  - Remove redundant parts of the conditions (e.g. `A || !A && B` can be `A || B`).

[logical_operators]: https://golang.org/ref/spec#Logical_operators
[if_statement]: https://golang.org/ref/spec#If_statements
[switch_statement]: https://golang.org/ref/spec#Switch_statements
