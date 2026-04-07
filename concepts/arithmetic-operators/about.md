# About Arithmetic Operators

## Operators

Go supports the following arithmetic operators:

| Operator | Example        |
| -------- | -------------- |
| `+`      | `4 + 6 == 10`  |
| `-`      | `15 - 10 == 5` |
| `*`      | `2 * 3 == 6`   |
| `/`      | `13 / 3 == 4`  |
| `%`      | `13 % 3 == 1`  |

For integer division, the remainder is dropped (e.g. `5 / 2 == 2`).

The `+` operator also concatenates strings: `"foo" + "bar" == "foobar"`.

## Mixed-Type Arithmetic

Go does not allow arithmetic operations between different numeric types:

```go
x := 42
n := 2.0
value := n * x // invalid operation: n * x (mismatched types float64 and int)
```

The operands must be the same type, sometimes requiring explicit conversion:

```go
x := float64(42)
n := 2.0
value := n * x 
```

## Shorthand Assignments

Each operator has a shorthand assignment form that updates the variable in place:

```go
a := 1
a += 2 // same as a = a + 2 == 3

b := 2
b -= 2 // same as b = b - 2 == 0

c := 4
c *= 2 // same as c = c * 2 == 8

d := 8
d /= 2 // same as d = d / 2 == 4

e := 16
e %= 2 // same as e = e % 2 == 0
```

## Increment and Decrement

When placed after a variable, the increment (`++`) and decrement (`--`) statements either increase or decrease its stored numeric value by 1.
Placing them before a variable is a syntax error.

```go
a := 10
a++ // same as a += 1, a == 11

b := 10
b-- // same as b -= 1, b == 9
```

Because they are statements, they do not return a value.
