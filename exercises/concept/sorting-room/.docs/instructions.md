# Instructions

Jen is working in the sorting room in a large factory.
The sorting room needs to process anything that comes into it by categorizing it with a label.
She needs a program to help her with the sorting.

Most primitive values should get straight-forward labels.
For numbers, she wants strings saying `"This is the number 2.0"` (if the number was 2).
Jen wants the same output for integers and floats.

There are a few `Box` interfaces that need to be unwrapped to get their contents.
For a `NumberBox`, she wants strings saying `"This is a box containing the number 3.0"` (if the `Number()` method returns 3).

Anything unexpected should say `"Return to sender"` so Jen can send them back where they came from.

## 1. Implement `DescribeNumber`

Jen wants numbers to return strings like `"This is the number 2.0"` (including one digit after the decimal):

```go
fmt.Println(DescribeNumber(-12.345))
// Output: This is the number -12.3
```

## 3. Implement `DescribeNumberBox`

Jen wants numbers to return strings like `"This is a box containing the number 2.0"` (again, including one digit after the decimal):

```go
fmt.Println(DescribeNumberBox(numberBoxContaining{12.345}))
// Output: This is a box containing the number 12.3
```

## 5. Implement `DescribeAnything` which uses them all

This is the main function Jen needs which takes any input (the empty interface means any value at all: `interface{}`).
`DescribeAnything` should delegate to the other functions based on the type of the value passed in.
More specifically:

- `int` and `float64` should both delegate to `DescribeNumber`
- `NumberBox` should delegate to `DescribeNumberBox`
- anything else should result in `"Return to sender"`

```go
fmt.Println(DescribeAnything(numberBoxContaining{12.345}))
// Output: This is a box containing the number 12.3
fmt.Println(DescribeAnything("some string"))
// Output: Return to sender
```
