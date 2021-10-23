# Introduction

Booleans in Go are represented by the `bool` type.
A `bool` is either `true` or `false`.

Go supports three boolean operators: `!` (NOT), `&&` (AND), and `||` (OR).

```go
true || false // Output: true
true && false // Output: false
!true // Output: false
```

The three boolean operators each have a different _operator precedence_.
As a consequence, they are evaluated in this order: `!` first, `&&` second, and finally `||`.
If you want to force a different ordering, you can enclose a boolean expression in parentheses (ie. `()`), as the parentheses have an even higher operator precedence.

```go
!true && false   // Output: false
!(true && false) // Output: true
```
