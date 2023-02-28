# Slice indexes

```go
// Package secret is a small library for interpreting signals for a secret greeting
package secret

// Handshake interprets what sequence of signals are contained within the number passed in
func Handshake(code uint) []string {
	signs := []string{"wink", "double blink", "close your eyes", "jump"}
	const reverseSigns = 16
	action, actionIncr, end := 0, 1, len(signs)

	if (code & reverseSigns) != 0 {
		action = 3
		actionIncr = -1
		end = -1
	}
	output := []string{}

	for ; action != end; action += actionIncr {
		if (code & (1 << action)) != 0 {
			output = append(output, signs[action])
		}
	}
	return output
}
```

This approach starts by defining a [slice][slice] of the signals in their normal order.
A [const] is defined to represent the signal for reversing the order of the other signals.
It is given a meaningful name so that `16` won't be used as a [magic number][magic-number].

Some variables are defined with values for iterating the slice of signals in their normal order.

The [bitwise AND operator][bitwise-operators] is used to check if the input number contains the signal for reversing the order of the other signals.

For example, if the number passed in is `19`, which is `10011` in binary, then it is ANDed with `16`, which is `10000` in binary.
The `1` in `10000` is also at the same position in `10011`, so the two values ANDed will not be `0`.
- `10011` AND
- `10000` =
- `10000`

If the number passed in is `3`, which is `00011` in binary, then it is ANDed with `16`, which is `10000` in binary.
The `1` in `10000` is not at the same position in `00011`, so the two values ANDed will be `0`.
- `00011` AND
- `10000` =
- `00000`

If the number passed in contains the signal for reverse, then the iteration variables are set to iterate backwards through the slice of signals.

The output slice is defined, and then the [`for` loop][for-loop] begins.

```go
for ; action != end; action += actionIncr {
```

Normal iteration will start at index `0`.
Reverse iteration will start at index `3`.

Normal iteration will terminate when the index equals `3`.
Reverse iteration will terminate when the index equals `-1`.

Normal iteration will increase the index by `1` for each iteration.
Reverse iteration will decrease the index by `1` for each iteration.

For each iteration, the AND operator is used to check if the number passed in contains `1` shifted left (`<<`) for the number of positions
as the value being iterated.

```go
if (code & (1 << action)) != 0 {
  output = append(output, signs[action])
}
```

For example, if the number being iterated is `0`, then `1` is shifted left `0` times (so not shifted at all), and the number passed in is ANDed with `00001`.
If the number passed in is `3`, which is `00011` in binary, then it is ANDed with `00001`.
`00011` ANDed with `00001` is not equal to `0`, so the signal at the index of the slice of signals is added to the output slice.
The index used is the number being iterated, which is `0`, so the element at index `0` ("wink`) would be added to the output slice
using the [append][append] function.

If the number being iterated is `1`, then `1` is shifted left `1` time, and the number passed in is ANDed with `00010`.
If the number passed in is `3`, which is `00011` in binary, then it is ANDed with `00010`.
`00011` ANDed with `00010` is not equal to `0`, so the signal at the index of the slice of signals is added to the output slice.
The index used is the number being iterated, which is `1`, so the element at index `1` ("double blink`) would be added to the output slice.

If the number passed in ANDed with the number being iterated is equal to `0`, then the signal in the slice for that index is not added to the output slice.

After iterating through the slice of signals is done, the output slice is returned from the function.

[slice]: https://go.dev/tour/moretypes/7
[const]: https://go.dev/tour/basics/15
[magic-number]: https://en.wikipedia.org/wiki/Magic_number_(programming)
[bitwise-operators]: https://www.tutorialspoint.com/go/go_bitwise_operators.htm
[for-loop]: https://go.dev/tour/flowcontrol/1
[append]: https://go.dev/tour/moretypes/15
