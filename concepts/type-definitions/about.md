# About

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

## Defined Types Are Distinct

Although `Score` and `int` both hold integer values, converting between them requires an explicit conversion:

```go
type Score int

var number int = 42

// var score Score = number is a compile error
var score Score = Score(number)
var converted int = int(score)
```

`42` is an untyped constant, so it can be assigned directly to a `Score` variable.

Defined types can also have methods, which are covered in the Methods concept.

## Struct Types

Struct definitions use the same `type` keyword:

```go
type Player struct {
    Name  string
    Score Score
}
```

Struct types are covered in the Structs concept.

## Type Aliases

A type alias creates another name for an existing type rather than creating a new type.
Because an alias and its original type are identical, their values can be used interchangeably.
Aliases are declared using the `type` keyword, the alias name, the equals sign (`=`), and the existing type.

```go
type Square = bool

var isOccupied bool = true
var square Square = isOccupied // no conversion needed
```
