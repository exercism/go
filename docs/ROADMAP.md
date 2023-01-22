# Roadmap

This document provides an overview of the main topics that contributors and maintainers will probably focus on in the coming months.

## Topics Overview

- Improve existing concept exercises
- Create additional concept exercises (see syllabus overview below)
- Define the `prerequisites` and `practices` lists for all practice exercises (see [Issue 1396](https://github.com/exercism/go/issues/1396))
- Fix the remaining [test runner](https://github.com/exercism/go-test-runner/) bugs
- Fix and reactivate the [analyzer](https://github.com/exercism/go-analyzer/) (e.g. see [Issue 16](https://github.com/exercism/go-analyzer/issues/16)) and fix other bugs
- Finish a good, stable version of the [representer](https://github.com/exercism/go-representer/) and take it live
- Improve/add approaches to exercises
- Sync tests with [problem specs](https://github.com/exercism/problem-specifications)

## Syllabus Overview

### Concept that exists

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
- `variadic-functions` (introduces `...`)
- `randomness`
- `regular-expressions`
- `stringers`

### Potential Future Concepts

- `time-duration`
- `integer-types`
- `panic-recover`
- `error-wrapping` (errors.Is, errors.As)
- `bytes`
- `bitwise-operators`
- `goroutines`
- `waitgroups`
- `mutexes`
- `channels` and `select`
- `closures`/`anonymous-functions`/`higher-order-functions`
- `JSON`
- `defer`
- `Buffer`, `io.Reader`/`io.Writer`
- `typed-parameters` (generics)
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
