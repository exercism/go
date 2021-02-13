# Implementation

The definition of the Cipher interface is located in
[cipher.go](./cipher.go).

Your implementations should conform to the Cipher interface.

```go
type Cipher interface {
    Encode(string) string
    Decode(string) string
}
```

It is expected that `Encode` will ignore all values in the string that
are not A-Za-z, they will not be represented in the output. The output
will be also normalized to lowercase.

The functions used to obtain the ciphers are:

```go
func NewCaesar() Cipher { }

func NewShift(distance int) Cipher { }

func NewVigenere(key string) Cipher { }
```

Argument for `NewShift` must be in the range 1 to 25 or -1 to -25.
Zero is disallowed.  For invalid arguments `NewShift` returns nil.

Argument for `NewVigenere` must consist of lower case letters a-z
only.  Values consisting entirely of the letter 'a' are disallowed.
For invalid arguments `NewVigenere` returns nil.

