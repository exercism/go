# Instructions

Jen is working in the sorting room in a large factory.
The sorting room needs to process anything that comes into it by categorizing it with a label.
She needs a program to help her with the sorting.

Most primitive values should get straight-forward labels.
For numbers, she wants strings saying `"This is the number 2.0"` (if the number was 2).
Jen wants the same output for integers and floats.
For booleans, she wants strings saying `"This is false"` or `"This is true"`.

There are a few `Box` interfaces that need to be unwrapped to get their contents.
For a `NumberBox`, she wants strings saying `"This is a box containing the number 3.0"` (if the `Number()` method returns 3).
For a `BooleanBox`, she wants strings saying `"This box is true"` or `"This box is false"` (depending on the result from the `Boolean()` method).

Anything unexpected should say `"Return to sender"` so Jen can send them back where they came from.

## 1. TODO

TODO:

```go
fmt.Println(CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake))
// Output: false
```

## 2. TODO

TODO:

```go
fmt.Println(CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake))
// Output: true
```
