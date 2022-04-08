# About 

Package [regexp][package-regexp] offers support for regular expressions in Go.

## Syntax 

The [syntax][regexp-syntax] of the regular expressions accepted is the same general syntax used by Perl, Python, and other languages. 

Both the search patterns and the input texts are interpreted as UTF-8.  

When using backticks (\`) to make strings, backslashes (`\`)  don't have any special meaning and don't mark the beginning of special characters like tabs `\t` or newlines `\n`:

```go
"\t\n" // regular string literal with 2 characters: a tab and a newline
`\t\n`// raw string literal with 4 characters: two backslashes, a 't', and an 'n'
```

Because of this, using backticks is desirable to make regular expressions,
because it means we don't need to escape backslashes:

```go
"\\" // string with a single backslash
`\\` // string with 2 backslashes
```

## Compiling patterns - `RegExp` type

The package defines a type `Regexp` for compiled regular expressions.
Function `regexp.Compile` compiles a string into a regular expression.
The function returns `nil` and an error if compilation failed:

```go
re, err := regexp.Compile(`(a|b)+`)
fmt.Println(re, err) // => (a|b)+ <nil>
re, err = regexp.Compile(`a|b)+`)
fmt.Println(re, err) // => <nil> error parsing regexp: unexpected ): `a|b)+`
```

Function `MustCompile` is a convenient alternative to `Compile`: 

```go 
re = regexp.MustCompile(`[a-z]+\d*`)
```

Using this function, there is no need to handle an error. 

~~~~exercism/caution
 `MustCompile` should only be used when we know for sure the pattern does compile, as otherwise the program will panic.
 ~~~~
 
 ## Regular expression methods
 
There are 16 methods of `Regexp` that match a regular expression and identify the matched text.
Their names are matched by this regular expression:

```text
Find(All)?(String)?(Submatch)?(Index)?
```

* If `All` is present, the routine matches successive non-overlapping matches of the entire expressions.
* If `String` is present, the argument is a string; otherwise it is a slice of bytes; return values are adjusted as appropriate. 
* If `Submatch` is present, the return value is a slice identifying the successive submatches of the expression.
* If `Index` is present, matches and submatches are identified by byte index pairs within the input string.

There are also methods for:

* replacing matches of regular expressions with replacement strings and
* splitting of strings separated by regular expressions.

All-in-all, the `regexp` package defines more than 40 functions and methods.
We will demonstrate the use of a few methods below.
Please see the [API documentation][package-regexp] for details of these and other functions.
 
### `MatchString` Examples 

Method `MatchString` reports whether a strings contains any match of a regular expression.

```go
re = regexp.MustCompile(`[a-z]+\d*`)
b = re.MatchString("[a12]")       // => true
b = re.MatchString("12abc34(ef)") // => true
b = re.MatchString(" abc!")       // => true
b = re.MatchString("123 456")     // => false    
```

### `FindString` Examples 

Method `FindString` returns a string holding the text of the leftmost match of the regular expression.

```go
re = regexp.MustCompile(`[a-z]+\d*`)
s = re.FindString("[a12]")       // => "a12"
s = re.FindString("12abc34(ef)") // => "abc34"
s = re.FindString(" abc!")       // => "abc"
s = re.FindString("123 456")     // => ""
```

### `FindStringSubmatch` Examples

Method `FindStringSubmatch` returns a slice of strings holding the text of the leftmost match of the regular expression and the matches, if any, of its subexpressions.
This can be used to identify the strings matching capturing groups.
A return value of `nil` indicates no match.

```go 
re = regexp.MustCompile(`[a-z]+(\d*)`)
sl = re.FindStringSubmatch("[a12]")       // => []string{"a12","12"}
sl = re.FindStringSubmatch("12abc34(ef)") // => []string{"abc34","34"}
sl = re.FindStringSubmatch(" abc!")       // => []string{"abc",""}
sl = re.FindStringSubmatch("123 456")     // => <nil>
```

Method `re.ReplaceAllString(src,repl)` returns a copy of `src`, replacing matches of the regular expression `re` with the replacement string `repl`.

```go
re = regexp.MustCompile(`[a-z]+\d*`)
s = re.ReplaceAllString("[a12]", "X")       // => "[X]"
s = re.ReplaceAllString("12abc34(ef)", "X") // => "12X(X)"
s = re.ReplaceAllString(" abc!", "X")       // => " X!"
s = re.ReplaceAllString("123 456", "X")     // => "123 456"
```
  
Method `re.Split(s,n)` slices a text `s` into substrings separated by the expression and returns a slice of the substrings between those expression matches.
The count `n` determines the maximal number of substrings to return.
If `n<0`, the method returns all substrings.

```go
re = regexp.MustCompile(`[a-z]+\d*`)
sl = re.Split("[a12]", -1)      // => []string{"[","]"}
sl = re.Split("12abc34(ef)", 2) // => []string{"12","(ef)"}
sl = re.Split(" abc!", -1)      // => []string{" ","!"}
sl = re.Split("123 456", -1)    // => []string{"123 456"}
```

## Performance

The regexp implementation provided by this package is guaranteed to run in 
[time linear in the size of the input][re2-performance].  

## Caveat
 
Package `regexp` implements [RE2 regular expressions][re2-syntax] (except for `\C`). 
The syntax is largely compatible with PCRE ("Perl Compatible Regular Expression"), but there are some differences.
Please see the "Caveat section" in [this article][reg-exp-wild] for details.
   
[raw-string-literals]:https://yourbasic.org/golang/regexp-cheat-sheet/#raw-strings
[package-regexp]:https://pkg.go.dev/regexp
[regexp-syntax]:https://pkg.go.dev/regexp/syntax
[re2-syntax]: https://golang.org/s/re2syntax
[reg-exp-wild]: https://swtch.com/~rsc/regexp/regexp3.html
[re2-performance]: https://swtch.com/~rsc/regexp/regexp1.html
