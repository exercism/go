Booleans in Go are represented by the `bool` type, which values can be either `true` or `false`.

Go supports three [boolean operators][logical operators]: `!` (NOT), `&&` (AND), and `||` (OR).

```go
true || false // => true
true && false // => false
!true // => false
```

The three boolean operators each have a different [_operator precedence_][operators]. As a consequence, they are evaluated in this order: `!` first, `&&` second, and finally `||`. If you want to 'escape' these rules, you can enclose a boolean expression in parentheses (`()`), as the parentheses have an even higher operator precedence.

```go
!true && false   // => false
!(true && false) // => true
```

[operators]: https://golang.org/ref/spec#Operators
[logical operators]: https://golang.org/ref/spec#Logical_operators
