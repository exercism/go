# Simple Cipher

Implement a simple shift cipher like Caesar and a more secure
substitution cipher.

## Step 1

"If he had anything confidential to say, he wrote it in cipher, that
is, by so changing the order of the letters of the alphabet, that not
a word could be made out. If anyone wishes to decipher these, and get
at their meaning, he must substitute the fourth letter of the
alphabet, namely D, for A, and so with the others."  —Suetonius, Life
of Julius Caesar

Ciphers are very straight-forward algorithms that allow us to render
text less readable while still allowing easy deciphering. They are
vulnerable to many forms of cryptoanalysis, but we are lucky that
generally our little sisters are not cryptoanalysts.

The Caesar Cipher was used for some messages from Julius Caesar that
were sent afield. Now Caesar knew that the cipher wasn't very good,
but he had one ally in that respect: almost nobody could read well. So
even being a couple letters off was sufficient so that people couldn't
recognize the few words that they did know.

Your task is to create a simple shift cipher like the Caesar Cipher.
This image is a great example of the Caesar Cipher:

![Caesar Cipher][1]

For example:

Giving "iamapandabear" as input to the encode function returns the
cipher "ldpdsdqgdehdu". Obscure enough to keep our message secret in
transit.

When "ldpdsdqgdehdu" is put into the decode function it would return
the original "iamapandabear" letting your friend read your original
message.

Initially you will implement a [Caesar Cipher][cc] with a fixed shift
distance of 3 (*namely D, for A*).

## Step 2

Fixed distance Shift Ciphers are no fun though when your kid sister
figures it out. Try amending the code to allow us to specify a shift
distance.

You will implement a more generic Shift Cipher with a flexible shift
distance.

# Step 3

With only 26 true possible shift values, your kid sister will figure
this out too. Next lets define a more complex cipher using a string as
key value: a [Vigenère cipher][vc].

Here's an example:

Given the key "aaaaaaaaaaaaaaaaaa", encoding the string
"iamapandabear" would return the original "iamapandabear".

Given the key "ddddddddddddddddd", encoding our string "iamapandabear"
would return the obscured "ldpdsdqgdehdu"

In the example above, we've set a = 0 for all the key values, with a
shift distance of 0. So when the plaintext is added to the key, we end
up with the same message coming out. So "aaaa" is not an ideal
key. But if we set the key to "dddd", we would get the same thing as
the Shift Cipher with all shift distances set to 4.

These keys are not much of an improvement over the fixed distance Shift
Cipher. However, we can put many different lengths into the key if we
use strings with different characters:

Given the key "adadadadad", encoding the string "iamapandabear" would
return the obscured "idmdpdngaeedr".

Each character in the key is used to shift the corresponding character
by index. If the key is shorter than the text, repeat the key as
needed.

## Extensions

Shift ciphers work by making the text slightly odd, but are vulnerable
to frequency analysis. Substitution ciphers help that, but are still
very vulnerable when the key is short or if spaces are
preserved. You'll see one solution to this problem in the exercise
"[crypto-square](https://github.com/exercism/go/tree/master/exercises/crypto-square)".

If you want to go farther in this field, the questions begin to be
about how we can exchange keys in a secure way. Take a look at
[Diffie-Hellman on Wikipedia][dh] for one of the first implementations
of this scheme.

[1]: https://upload.wikimedia.org/wikipedia/commons/thumb/4/4a/Caesar_cipher_left_shift_of_3.svg/320px-Caesar_cipher_left_shift_of_3.svg.png
[cc]: https://en.wikipedia.org/wiki/Caesar_cipher
[vc]: https://en.wikipedia.org/wiki/Vigen%C3%A8re_cipher
[dh]: https://en.wikipedia.org/wiki/Diffie%E2%80%93Hellman_key_exchange

## Implementation

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



## Running the tests

To run the tests run the command `go test` from within the exercise directory.

If the test suite contains benchmarks, you can run these with the `--bench` and `--benchmem`
flags:

    go test -v --bench . --benchmem

Keep in mind that each reviewer will run benchmarks on a different machine, with
different specs, so the results from these benchmark tests may vary.

## Further information

For more detailed information about the Go track, including how to get help if
you're having trouble, please visit the exercism.io [Go language page](http://exercism.io/languages/go/resources).

## Source

Substitution Cipher at Wikipedia [http://en.wikipedia.org/wiki/Substitution_cipher](http://en.wikipedia.org/wiki/Substitution_cipher)

## Submitting Incomplete Solutions
It's possible to submit an incomplete solution so you can see how others have completed the exercise.
