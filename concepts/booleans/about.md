# About

A Boolean value represents whether a condition is `true` or `false`.
In Go, the `bool` type has only those two values.

```go
isDoorLocked := true
hasKey := true
knowsCode := false
```

## Logical Operators

Logical operators work on one or two `bool` values and produce a Boolean result.

| Operator | Name | Behavior |
| -------- | ---- | -------- |
| `!`      | NOT  | Returns the opposite of a single Boolean value |
| `&&`     | AND  | Returns `true` only when both sides are `true` |
| `\|\|`     | OR   | Returns `true` if any side is `true` |

Use `!` (NOT) when you need the opposite of a Boolean value:

```go
!isDoorLocked // false
```

Use `&&` (AND) when both conditions must be `true`:

```go
canUnlockDoor := isDoorLocked && hasKey // true
```

Use `||` (OR) when at least one condition must be `true`:

```go
canOpenDoor := canUnlockDoor || knowsCode // true
```

Boolean operators are evaluated by default in this order: `!` first, then `&&`, then `||`.
Use parentheses to override this behavior:

```go
!true && false   // false
!(true && false) // true
```
