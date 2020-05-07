### Strings

A `string` in Go is a sequence of `bytes`, which doesn't necessarily have to represent characters.
That being said, `UTF-8` is a central part of `strings` in Go. It is easy to convert a string to `runes` (`UTF-8` characters) or iterate over `runes` in a string.
This makes dealing with different languages or other special characters in Go very simple.

When dealing with `UTF-8` characters it is important to know that not all characters have the same length;
`ASCII` characters have a length of one `byte`, while other characters can be as long as four `bytes`.
Runes, bytes, and their connection to strings will be handled more in-depth in a later exercise.

The underlying type (aka representation) for `string` in Go is the `[]byte` slice, which allows for some commonalities with slices in general.
For example, you can get a `byte` at position `i` from a `string` called `s` with sub-indexing, like `s[i]`.
With that said, string types are immutable so operations like s[i] = 'a' are not available and will result in a compilation error.

[Strings, bytes, runes and characters in Go](https://blog.golang.org/strings) provides a deep dive into this topic.

### Conditionals

This exercise also introduces _conditionals_. Here is a little intro:
[Go by Example: If/Else](https://gobyexample.com/if-else)
