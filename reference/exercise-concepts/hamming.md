# Hamming

## Exercise

Calculate the Hamming Distance between two DNA strands.

## Concepts encountered while solving it

## Strings

Strings in Go are made up of runes. Runes are Go's equivalent to the concept of characters. A rune can be represented using from 1 up to 4 bytes.

### String iteration

- When iterating over strings using the range operator, runes are yielded. When dealing with ASCII runes, strings can be safely ranged using the string[i] syntax. However when non ASCII characters are present this is unsafe as only one byte of the corresponing rune is returned.

### String conversion

- Since two strings cannot be ranged using the range operator at the same time, the safe approach for comparing each "character" in two strings is by first casting the strings to a slice of runes.

## Errors

### Error Creation

- In Go there are no exceptions, instead errors can be defined and processed accordingly to the context. Errors can be defined using fmt.Errorf or errors.New. The latter is preferred when there is no string interpolation within the error message. Error messages should be concise, not vague and provide context for where and why the error occurred. Also they should not be capitalized or have punctuation.

### Returning Errors

- When an error is encountered the convention is to return a zero value along with the corresponging error. In the opposite case a nil error should be returned along with the value that the function should return.

Note: An error should be the last returned parameter of any given function.

## Other "Minor" Concepts

### Multiple parameters

A function can have many parameters. Go has a handy shortcut for same-type parameters allowing to specify the type only once after the parameter names.

### Multiple return values

Go allows multiple return values. This can be leveraged in many ways but it is most often used for returning errors alongside other values.
