# Hints

## General

- There are three [boolean operators][logical operators] to work with boolean values.
- Multiple operators can be combined in a single expression.

## 1. Check if a fast attack can be made

- The logical NOT operator (`!`) can be placed before an expression to negate its value.

## 2. Check if a spy action can be made

- Logical operators are typically used to evaluate whether two or more expressions are true or not true.

## 3. Check if a signal action can be made

- Logical operators execute in the order of their precedence (from highest to lowest): `!`, `&&`, `||`.
- For more details check out the Operator Precedence section on the [official golang documentation][operators] and the [truth table][truth table].

[logical operators]: https://golang.org/ref/spec#Logical_operators
[operators]: https://golang.org/ref/spec#Operators
[truth table]: https://www.digitalocean.com/community/tutorials/understanding-boolean-logic-in-go
