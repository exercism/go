# Introduction

A type definition creates a new named type, distinct from every other type, even one with the same underlying type.

## Defining Types

A type definition starts with the `type` keyword, followed by the new type's name and its underlying type:

```go
type File []bool
```

`File` here refers to a single vertical column on a chessboard, commonly labeled A to H.
Each square in this column is represented by a Boolean value marking whether it's occupied or not.
Now that `File` is defined, we can use it inside other definitions:

```go
type Chessboard map[string]File

board := Chessboard{
    "A": File{true, false, true},
    "B": File{false, true, false},
}
```

`Chessboard` maps file identifiers to `File` values.
Each `File` is a slice of Booleans representing the squares in one chess file.
Without these type definitions, `board` would be a `map[string][]bool`, which would make it harder to understand what we're representing.
