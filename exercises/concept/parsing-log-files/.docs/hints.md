# General

- API documentation: [package regexp][package-regexp]
- Regular expression syntax: [regexp/syntax][regexp-syntax]
- Website [regex101.com][regex101] has an online regular expression tester.
- It is recommended to write regular expressions as [raw string literals][raw-string-literals] enclosed by backticks.

## 1. Identify garbled log lines

- Function [regexp.MatchString][fun-match-string] or method [MatchString][method-match-string] could be useful here.

## 2. Split the log line

- Method [Split][regexp-split] could be useful here.
  
## 3. Count the number of lines containing `password` in quoted text
 
- You can make expression matching case sensitive by prefixing the regular expression with `(?i)`.
This will set the `i` flag.  See [this tutorial][yourbasic-i-flag].

## 4. Remove artifacts from log

- Method [ReplaceAllString][replace-all-string] could be useful here.

## 5. Tag lines with user names
 
- Method [FindStringSubmatch][find-string-submatch] could be useful here.

[raw-string-literals]: https://yourbasic.org/golang/regexp-cheat-sheet/#raw-strings
[package-regexp]: https://pkg.go.dev/regexp
[regexp-syntax]: https://pkg.go.dev/regexp/syntax
[regex101]: https://regex101.com/
[fun-re-match-string]: https://pkg.go.dev/regexp#MatchString
[method-match-string]: https://pkg.go.dev/regexp#Regexp.MatchString
[regexp-split]: https://pkg.go.dev/regexp#Regexp.Split
[yourbasic-i-flag]: https://yourbasic.org/golang/regexp-cheat-sheet/#case-insensitive-and-multiline-matches
[replace-all-string]: https://pkg.go.dev/regexp#Regexp.ReplaceAllString
[find-string-submatch]: https://pkg.go.dev/regexp#Regexp.FindStringSubmatch