# Introduction

The `rune` type in Go is an alias for `int32`. Given this underlying `int32` type, the `rune` type holds a signed 32-bit integer value. However, unlike an `int32` type, the integer value stored in a `rune` type is meant to represent a single Unicode character.

## Unicode and Unicode Code Points

Unicode is a superset of ASCII that represents characters by assigning a unique number to every character. This unique number is called a Unicode code point. Unicode aims to represent all the world's characters including various alphabets, numbers, symbols, and even emoji as Unicode code points.

In Go, the `rune` type is meant to represent a single Unicode code point.

The following table contains example Unicode characters along with their Unicode code point and decimal values:

| Unicode Character | Unicode Code Point | Decimal Value |
|-------------------|--------------------|---------------|
| 0                 | `U+0030`           | `48`          |
| A                 | `U+0041`           | `65`          |
| a                 | `U+0061`           | `97`          |
| ¿                 | `U+00BF`           | `191`         |
| π                 | `U+03C0`           | `960`         |
| 🧠                | `U+1F9E0`          | `129504`      |

## UTF-8

UTF-8 is a variable-width character encoding that is used to encode every Unicode code point as 1, 2, 3, or 4 bytes. Since a Unicode code point can be encoded as a maximum of 4 bytes, the `rune` type needs to be able to hold up to 4 bytes of data. That is why the `rune` type is an alias for `int32` as an `int32` type is capable of holding up to 4 bytes of data.

Go source code files are encoded using UTF-8.

## Using Runes

Variables of type `rune` are declared by placing a character inside single quotes:

```go
myRune := '¿'
```

Since the underlying type of a rune is `int32`, printing a rune's type will yield `int32`:

```go
myRune := '¿'
fmt.Printf("myRune type: %T\n", myRune)
// Output: myRune type: int32
```

Similarly, printing a rune's value will yield the underlying integer (decimal) value:

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

Since Go source code files are encoded using UTF-8, strings in Go are also encoded using UTF-8. That means strings in Go contain Unicode characters. Since a `rune` type is meant to represent a Unicode character, a string in Go is often referred to as a sequence of runes. Taking this one step further, since runes are stored as 1, 2, 3, or 4 bytes, a string can also be referred to as a sequence of bytes. Go uses slices to represent such sequences, and these slices can be iterated over using `range`.

To iterate over the runes of a given string, use `range`. In this example, the `index` variable represents the starting index of the current rune's byte sequence and the `char` variable represents the current rune:

```go
myString := "exercism"
for index, char := range myString {
  fmt.Printf("Index: %d\tCharacter: %c\t\tCode Point: %U\n", index, char, char)
}
// Output:
// Index: 0	Character: e		Code Point: U+0065
// Index: 1	Character: x		Code Point: U+0078
// Index: 2	Character: e		Code Point: U+0065
// Index: 3	Character: r		Code Point: U+0072
// Index: 4	Character: c		Code Point: U+0063
// Index: 5	Character: i		Code Point: U+0069
// Index: 6	Character: s		Code Point: U+0073
// Index: 7	Character: m		Code Point: U+006D
```

Since runes can be stored as 1, 2, 3, or 4 bytes, the length of a string may not always equal the number of characters in the string:

```go
myString1 := "exercism"
myString2 := "🧠exercism"
fmt.Printf("myString1 length: %d\n", len(myString1))
fmt.Printf("myString2 length: %d\n", len(myString2))
// Output:
// myString1 length: 8
// myString2 length: 12
```

## Type Converting Runes

A slice of runes can be type converted to a string:

```go
myRuneSlice := []rune{'e', 'x', 'e', 'r', 'c', 'i', 's', 'm'}
myString := string(myRuneSlice)
fmt.Println(myString)
// Output: exercism
```

Similarly, a string can be type converted to a slice of runes. Remember, without formatting verbs, printing a rune yields its underlying integer value:

```go
myString := "exercism"
myRuneSlice := []rune(myString)
fmt.Println(myRuneSlice)
// Output: [101 120 101 114 99 105 115 109]
```
