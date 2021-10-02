# Introduction

The `rune` type in Go is an alias for `int32`. Given this underlying `int32` type, the `rune` type holds a signed 32-bit integer value. However, unlike an `int32` type, the integer value stored in a `rune` type is represents a single Unicode character.

## Using Runes

Variables of type `rune` are declared by placing a character inside single quotes:

```go
myRune := '¿'
```

Since `rune` is just an alias for `int32`, printing a rune's type will yield `int32`:

```go
myRune := '¿'
fmt.Printf("myRune type: %T\n", myRune)
// Output: myRune type: int32
```

Similarly, printing a rune's value will yield its integer (decimal) value:

```go
myRune := '¿'
fmt.Printf("myRune value: %v\n", myRune)
// Output: myRune value: 191
```

To print the Unicode character represented by the rune, use the `%c` formatting verb:

```go
myRune := '¿'
fmt.Printf("myRune Unicode character: %c\n", myRune)
// Output: myRune Unicode character: ¿
```

To print the Unicode code point represented by the rune, use the `%U` formatting verb:

```go
myRune := '¿'
fmt.Printf("myRune Unicdoe code point: %U\n", myRune)
// Output: myRune Unicode code point: U+00BF
```

## Runes and Strings

Strings in Go are encoded using UTF-8 which means they contain Unicode characters. Since the `rune` type represents a Unicode character, a string in Go is often referred to as a sequence of runes. However, runes are stored as 1, 2, 3, or 4 bytes depending on the character. Due to this, strings are really just a sequence of bytes. In Go, slices are used to represent sequences and these slices can be iterated over using `range`.

Even though a string is just a slice of bytes, the `range` keyword iterates over a string's runes, not its bytes. In this example, the `index` variable represents the starting index of the current rune's byte sequence and the `char` variable represents the current rune:

```go
myString := "❗hello"
for index, char := range myString {
  fmt.Printf("Index: %d\tCharacter: %c\t\tCode Point: %U\n", index, char, char)
}
// Output:
// Index: 0	Character: ❗		Code Point: U+2757
// Index: 3	Character: h		Code Point: U+0068
// Index: 4	Character: e		Code Point: U+0065
// Index: 5	Character: l		Code Point: U+006C
// Index: 6	Character: l		Code Point: U+006C
// Index: 7	Character: o		Code Point: U+006F
```

Since runes can be stored as 1, 2, 3, or 4 bytes, the length of a string may not always equal the number of characters in the string:

```go
myString1 := "hello"
myString2 := "❗hello"
fmt.Printf("myString1 length: %d\n", len(myString1))
fmt.Printf("myString2 length: %d\n", len(myString2))
// Output:
// myString1 length: 5
// myString2 length: 8
```
