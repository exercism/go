# Hints

## General

Conditionals are used to check for certain conditions and/or criteria. The most basic way of performing a conditional operation is using a single `if` statement.

## 1. Multiple cases (else if)

For various cases an `else if` and `else` statements can be used like this:

```go
if condition {
    // some code
} else if other_condition {
    // some code
} else if another_condition {
    // some code
} else {
    // some code
}
```

The `else` statement is any other case not covered by previous conditions. It should be noted that since Go allows early returns it is common to avoid the `else` statement after a return.

## 2. Multiple cases (logical operators)

Another way of checking for different scenarios is by chaining logical operators within the `if` statements like this:

```go
if conditionA ||conditionB {
    // some code
} else if conditionC && conditionD {
    // some code
} else {
    // some code
}
```

## 3. More than 3 cases (switch)

When there are 3 or more cases it is often better to use a `switch` statement, which allows to test for multiple conditions. It also has a `default` for managing unmanaged cases. It is used like this:

```go
switch {
case aCase || bCase:
    // some code
case otherCase:
    // some code
case anotherCase:
    // some code
default:
    // some code
}
```

Note that within each case different logical operators can be chained together to expand the evaluation.
