# About

In Go, we can compare values using the equality and ordering operators.
We have 6 of these operators:

| Comparison        | Operator  |
| ------------------| --------- |
| equal             | `==`      |
| not equal         | `!=`      |
| less              | `<`       |
| less or equal     | `<=`      |
| greater           | `>`       |
| greater or equal  | `>=`      |

The result of a comparison is always a boolean value:

```go
a := 3

a != 4 // true
a > 5  // false
```

## Equality operators

The equality operators `==` and `!=` check whether a value is equal to another or not, respectively.
We can use the equality operators on the majority of types in Go. Here are some common examples:

```go
2 == 3 // false, integer comparison

2.1 != 2.2 // true, float comparison 

"hello" == "hello" // true, string comparison

'a' != 'b' // true, rune/byte comparison
```

## Ordering operators

Ordering operators check if one value is greater than (`>`), greater or equal to (`>=`), less than (`<`) and less or equal to (`<=`) to another value.
This kind of comparison is available for numbers, bytes, runes and strings.
When comparing strings, the dictionary order (also known as lexicographic order) is followed.

```go
2 > 3 // false, integer comparison

1.2 < 1.3 // true, float comparison

'a' > 'b' // false, rune/byte comparison

"Hello" < "World" // true, string comparison
```
