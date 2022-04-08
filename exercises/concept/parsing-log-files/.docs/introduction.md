# Introduction

Package [regexp][package-regexp] offers support for regular expressions in Go.
The [syntax][regexp-syntax] of the regular expressions accepted is the same general syntax used by Perl, Python, and other languages.

The package defines a type `Regexp` for compiled regular expressions.
Function `regexp.Compile` compiles a string into a regular expression.
The function returns `nil` and an error if compilation failed:

```go
re, err := regexp.Compile(`(a|b)+`)
fmt.Println(re, err) // => (a|b)+ <nil>
re, err = regexp.Compile(`a|b)+`)
fmt.Println(re, err) // => <nil> error parsing regexp: unexpected ): `a|b)+`
```

When using backticks ``(`)`` to make strings, backslashes `(\)` don't have any special meaning, and don't mark the beginning of special characters like tabs `\t` or
newlines `\`n:

```go
"\t\n" // regular string literal with 2 characters: a tab and a newline
`\t\n`// raw string literal with 4 characters: two backslashes, a 't', and an 'n'
```

Because of this, using backticks is desirable to make regular expressions, because it means we don't need to escape backslashes:

```go
"\\" // string with a single backslash
`\\` // string with 2 backslashes
```
  
The `regexp` package defines more than 40 functions and methods.
We will demonstrate the use of a few methods below.
Please see the [API documentation][package-regexp] for details of these and other functions.

Method `MatchString` reports whether a strings contains any match of a regular expression.

```go
re,_ = regexp.Compile(`[a-z]+\d*`)
b = re.MatchString("[a12]")       // => true
b = re.MatchString("12abc34(ef)") // => true
b = re.MatchString(" abc!")       // => true
b = re.MatchString("123 456")     // => false    
```
 
Method `FindString` returns a string holding the text of the leftmost match of the regular expression.

```go
re,_ = regexp.Compile(`[a-z]+\d*`)
s = re.FindString("[a12]")       // => "a12"
s = re.FindString("12abc34(ef)") // => "abc34"
s = re.FindString(" abc!")       // => "abc"
s = re.FindString("123 456")     // => ""
```

Method `FindStringSubmatch` returns a slice of strings holding the text of the leftmost match of the regular expression and the matches, if any, of its subexpressions.
This can be used to identify the strings matching capturing groups.
A return value of `nil` indicates no match.

```go 
re,_ = regexp.Compile(`[a-z]+(\d*)`)
sl = re.FindStringSubmatch("[a12]")       // => []string{"a12","12"}
sl = re.FindStringSubmatch("12abc34(ef)") // => []string{"abc34","34"}
sl = re.FindStringSubmatch(" abc!")       // => []string{"abc",""}
sl = re.FindStringSubmatch("123 456")     // => <nil>
```

Method `re.ReplaceAllString(src,repl)` returns a copy of `src`, replacing matches of the regular expression `re` with the replacement string `repl`.

```go
re,_ = regexp.Compile(`[a-z]+\d*`)
s = re.ReplaceAllString("[a12]", "X")       // => "[X]"
s = re.ReplaceAllString("12abc34(ef)", "X") // => "12X(X)"
s = re.ReplaceAllString(" abc!", "X")       // => " X!"
s = re.ReplaceAllString("123 456", "X")     // => "123 456"
```
  
Method `re.Split(s,n)` slices a text `s` into substrings separated by the expression and returns a slice of the substrings between those expression matches.
The count `n` determines the maximal number of substrings to return.
If `n<0`, the method returns all substrings.

```go
re,_ = regexp.Compile(`[a-z]+\d*`)
sl = re.Split("[a12]", -1)      // => []string{"[","]"}
sl = re.Split("12abc34(ef)", 2) // => []string{"12","(ef)"}
sl = re.Split(" abc!", -1)      // => []string{" ","!"}
sl = re.Split("123 456", -1)    // => []string{"123 456"}
```
  
[raw-string-literals]: https://yourbasic.org/golang/regexp-cheat-sheet/#raw-strings
[package-regexp]: https://pkg.go.dev/regexp
[regexp-syntax]: https://pkg.go.dev/regexp/syntax
