# Roadmap

This document provides an overview of the main topics that contributors and maintainers will probably focus on in the coming months.

## Topics Overview

- Improve existing concept exercises
- Create additional concept exercises (see syllabus overview below)
- Define the `prerequisites` and `practices` lists for all practice exercises (see [Issue 1396](https://github.com/exercism/go/issues/1396))
- Update to the latest Go version ([Issue 1881](https://github.com/exercism/go/issues/1881))
- Fix the remaining [test runner](https://github.com/exercism/go-test-runner/) bugs
- Implement a `golint` replacement in the [analyzer](https://github.com/exercism/go-analyzer/) ([Issue 16](https://github.com/exercism/go-analyzer/issues/16)) and fix other bugs
- Add analyzer comments for all existing concept exercises
- Add more mentor notes for practice exercises
- Improve/update/proof-read various documentation (contributing guidelines, docs, exercises/shared, reference folder)
- Sync tests with [problem specs](https://github.com/exercism/problem-specifications)

## Syllabus Overview

### Concept that exists or an issue has been created

- `basics`
- `strings` and `strings-package`
- `numbers` (basic) and `arithmetic-operators`
- `floating-point-numbers`
- `booleans`
- `comments`
- `conditionals-if` and `comparison`
- `conditionals-switch`
- `structs`
- `type-definitions`
- `slices`
- `iteration` (for loop)
- `functions` (incl. multiple return values)
- `maps`
- `range-iteration`
- `runes`
- `pointers`
- `methods`
- `interfaces`
- `errors`
- `zero-values` (incl. nil)
- `type-conversion`
- `type-assertion`
- `constants`
- `string-formatting`
- `packages` (import/export)
- `time`
- `time-duration`
- `integer-types`
- `panic-recover`

### Potential Future Concepts

- `bytes`
- `bitwise-operators`
- `goroutines`
- `waitgroups`
- `mutexes`
- `channels` and `select`
- `closures`/`anonymous-functions`/`higher-order-functions`
- `rest-and-spread` (introduces the `...` operator)
- `JSON`
- `defer`
- `randomness`
- `regular-expressions`
- `stringers`
- `Buffer`, `io.Reader`/`io.Writer`
- error wrapping (errors.Is, errors.As) and custom errors
- `generics` (once they become available)
- `context`
- time helpers (ticket, sleep, etc.)
- `reflection`
- `unsafe`
- `init`
- functional options pattern
- embedding files

Additional topics that might be out of scope for concept exercises:

- tooling (go fmt, go vet, go doc etc.)
- writing tests and examples
- benchmarking
- profiling
- go generate
- compiler options
