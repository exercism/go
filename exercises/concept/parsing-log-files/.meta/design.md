# Design

## Learning objectives

In this concept the student should learn about the following topics and then practice them in the concept exercise.

* which flavor of regex syntax is used in Go and its performance
* limitations of regex in Go (see ["Caveats" section][regex-caveats]).
* how to create a regular expression with `Compile`/`MustCompile` and when to use which
* how to use the most common regex related functions
* understanding difference between `MatchX` and `FindX`
* understanding that there are byte and string versions of the functions
* understanding the different parts of the `Submatch` results
* how to set flags
* using e.g. [regex101] to help with writing/reading regex

## Out of Scope

Explaining how to write regular expressions themselves is out of scope for the concept here but we should link to some good resource a student could read to learn about them from scratch.
We don't do this as part of the concept because Exercism assumes the student is already fluent in another language and most languages include some form of regular expressions.

## Prerequisites

* `runes` so the student understand the string vs bytes distinction
* `slices` to understand the result of FindX etc
* `methods` because most functionality is provided by methods defined for the `RegExp` type

## Exercise Idea

C# ["Parsing Log Lines"][c#-parsing-log-lines] Exercise could serve as template.

In case you port that exercise, make sure to check all tasks make sense for Go and whether some additional task needs to be added to cover some Go specific
content.

## Resources

Here some links that might be helpful as a starting point and/or for the links section of the concept:

* [Go regexp package][regexp-package]
* [Go regexp syntax][regexp-syntax]
* [re2 syntax][re2-syntax]
* [Go By Example: regexp][gobyexample]

[c#-parsing-log-lines]: https://github.com/exercism/csharp/tree/main/exercises/concept/parsing-log-files
[regexp-package]: https://pkg.go.dev/regexp
[regexp-syntax]: https://pkg.go.dev/regexp/syntax
[re2-syntax]: https://github.com/google/re2/wiki/Syntax
[gobyexample]: https://gobyexample.com/regular-expressions
[regex-caveats]: https://swtch.com/~rsc/regexp/regexp3.html
[regex101]: https://regex101.com/
