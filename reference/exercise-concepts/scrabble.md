# Concepts of Scrabble

There are different approaches to Scrabble which are in the different categories below. Common concepts for all approaches are in 'General'.

## General

- zero value variable declaration: The `score` variable for the total score needs initialization to `0` (idiomatically with `var score int`).
- strings: lower and upper case letters should be treated equally (`ToUpper`/`ToLower`).
- range loop over string: using `range` on a string is a bit special. It returns the byte index of each rune, and the rune value itself.
- for loop: This exercise can also be solved with a `for i := 0...` loop -- although `for range` is preferable here.
- runes: using a range loop over a string will return a rune.
- bytes: using a for loop comes with indexing into a string which returns a `byte`. In general, this is unsafe, unless you can guarantee that the string contains only ASCII runes. Prefer `range` on strings, or conversion to a `[]rune`.
- unicode: instead of the entire string (`strings`) a single `rune` can be transformed `ToUpper`/`ToLower`.
- numbers: adding a value to an existing sum. (`++`, `+=`)
- comments: comment on exported function should be present. Best also on package itself, but that is more optional.
- ignore values: In the `range` loop the index is not needed and has to be ignored with a `_` (underscore).

## Approach: Switch

- switch: uses switch to get the score per letter
- functions: makes the code nicer to extract the big switch, but the function call overhead adds about 30% to the benchmark time

## Approach: map

- map: is used to store the score per letter. Map lookup. Map definition with content.
- globals: The `map` containing the scores is effectively static and should not be defined on each call to the function. Making it a global is one approach to solve this.
- init function: some use an `init` function to initialize the map from something faster to type, but this is lazy; code should be optimised for the reader, not the writer.
- type conversion: some might use a `map[string]int` and need to convert from `rune` to `string`.

## Approach: Slice

Some take the `slice` approach as it is fastest. However, a map or switch is more natural, and at this stage of learning, we should prefer naturalness and simplicity.

- slice: a slice is used to store the scores in alphabetical order.
- rune maths: a rune is just a number. This is used to convert a rune into a slice index: `r - 'A'`
- if condition: if using a slice, an out of bounds index would `panic`. So best to check for unknown indexes.
- panic: what is a `panic`, how it can be `recover`ed. Fixing is the right approach for `panic`s, handling only for unknown bugs.
- errors: errors vs panics.
