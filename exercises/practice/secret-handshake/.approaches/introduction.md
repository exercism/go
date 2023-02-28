# Introduction

There are several variants of a general approach for solving Secret Handshake.
The general approach is to use slice indexes for comparing bitwise values.

## General guidance

Something to consider is to keep the number of iterations at a minimum to get the best performance.

## Approach: Slice indexes

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

For more information, check the [Slice indexes approach][approach-slice-indexes].

[approach-slice-indexes]: https://exercism.org/tracks/go/exercises/secret-handshake/approaches/slice-indexes
