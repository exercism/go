# Roadmap

This document provides an overview of the main topics that contributors and maintainers will probably focus on in the coming months.

## Topics Overview

- Improve existing concept exercises
- Create additional concept exercises (see syllabus overview below)
- Define the `prerequisites` and `practices` lists for all practice exercises (see [tracking issue](https://github.com/exercism/javascript/issues/960))
- Improve the [test runner](https://github.com/exercism/javascript-test-runner/)
  - Fix bugs
  - Nicer output in case variables are defined outside of the test case
  - Connect tests to tasks
- Add [analyzer](https://github.com/exercism/go-analyzer/) comments for all existing concept exercises
- Improve/update/proof-read various documentation (contributing guidelines, how to implement a concept exercise, docs, exercises/shared, reference folder)

## Syllabus Overview

### Concept that exists or an issue has been created

- `basics`
- `strings` (incl. string methods)
- `numbers`
- `booleans`
- `conditionals` (if/else) and `comparison`
- `for-loops` and `increment-decrement`
- `while-loops`
- `conditionals-switch`
- `objects` (as key-value collections)
- `null-undefined`
- `functions`
- `type-conversion` (incl. truthy/falsy)
- `template-strings`
- `conditionals-ternary`
- `callbacks`
- `arrow-functions`
- `classes` and `prototypes`
- `inheritance`
- `errors`
- `arrays` (incl. array manipulation with push, pop etc.)
- `array-loops`
- `array-analysis`
- `array-transformation`
- `array-destructering`
- `rest-and-spread`
- `promises`
- `recursion`
- `closures`
- `sets`

### Potential Future Concepts

- `type-checking`
- `maps` (should be added to the Sets exercise, WeakSet/WeakMap could be in about as first iteration)
- `dates`
- `async-await`
- `JSON`
- `randomness`
- `regular-expressions`
- `math`
- getters/setters, flags, descriptors
- `proxy`
- `variables` (deep dive on const/let/var)
- `immutability` (incl. freeze, could be combined with the variables concept)
- `object-helpers` (e.g. Object.assign)
- chars and codepoints
- this and scope (incl. call/bind)
- iterators and "enumerators"
- generators/yield
- asnyc generators and iterators
- bit wise operators
- modules, imports, exports, dynamic imports
- duck-typing
- Internationalization API
- `bigint`
- `symbol`
- JSDoc comments
- asychronicity, event loop
- setTimeout/setInterval
- events
- "everything is an object" sometimes (relation of functions, arrays and primitives to objects, e.g. boxed types)
