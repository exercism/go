# Zero Values

Every type in Go has a [zero value](https://golang.org/ref/spec#The_zero_value).
If a variable is declared and no value assigned, its zero value is automatically assigned to it.
A variable in Go therefore can never be empty or undefined, etc.

The zero values are:
- `false` for booleans
- `""` (empty string) for strings
- `0` for any numerical type (int, in64, float64, uint8, byte, etc.)
- `nil` for pointers, functions, interfaces, slices, channels and maps
