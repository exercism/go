# Generate sequentially and shuffle

```go
// Package robotname is a small library for generating robot names.
package robotname

import (
	"fmt"
	"math/rand"
)

// Robot is a struct with a string field.
type Robot struct {
	name string
}

const maxRobotNames = 26 * 26 * 10 * 10 * 10

var (
	namePool = generateRobotNames()
	idx      = 0
)

func generateRobotNames() []string {
	pos := 0
	names := make([]string, maxRobotNames)

	for i := 'A'; i <= 'Z'; i++ {
		for j := 'A'; j <= 'Z'; j++ {
			for k := 0; k < 1000; k++ {
				names[pos] = fmt.Sprintf("%s%s%03d", string(i), string(j), k)
				pos++
			}
		}
	}

	rand.Shuffle(len(names), func(i, j int) { names[i], names[j] = names[j], names[i] })
	return names
}

// Name returns the existing name or returns a newly assigned name.
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	if idx >= maxRobotNames {
		return "", fmt.Errorf("uniqueness exhausted")
	}

	r.name = namePool[idx]
	idx++

	return r.name, nil
}

// Reset erases the existing name without returning it to unused names.
func (r *Robot) Reset() {
	r.name = ""
}

```

The approach imports some packages that will be needed for randomization and formatting the name.

The `Robot` type is defined as a struct with a `name` field of type `string`.

A [constant][const] is defined for the maximum number of names.
Then the pool of names is defined along with the current index to be used for accessing the name pool.


The function for generating the robot names is where most of the work is done.
A variable is defined to keep track of the position in the name pool for each name as it is generated.
The [make][make] function is used to create the [slice][slice] for all of the names.

- The outer `for` loop generates the first letter for the name.
- The middle `for` loop  generates the second letter for the name.
- The inner `for` loop generates the number for the name.

[`fmt.Sprintf()`][sprintf] is used in the body of the inner `for` loop for formatting the current values of the letters and number into the name
for that position.
The format specifier `%03d` says to pad a base-10 number with zeros for up to three places.
If one place is occupied by a digit, then the two extra places are padded with two zeros.
More info can be found at a [format cheat sheet][format-cheat-sheet].

After the loops finish, the [`rand.Seed()`][seed] function is called to initialize the random [`Source`][source] of numbers.
It is done by passing the [current time][now] in [nanoseconds][unixnano].
The [`rand.Shuffle()`][shuffle] function is then called to shuffle the name pool, after which the name pool is returned from the `generateRobotNames()` function.

When the `Name()` method is called, it first checks if a name has already been assigned and returns it without error if so.
If not, it checks if all possible names have been used.
If all possible names have been used, then the function returns an error.
Otherwise, the currently available name is assigned to the `Robot`, the index into the name pool is incremented,
and the newly assigned name is returned without error.

The `Reset()` method ensures the `Robot` name is set to its [zero value][zero-value] of an empty string.

[const]: https://go.dev/tour/basics/15
[make]: https://go.dev/tour/moretypes/13
[slice]: https://go.dev/tour/moretypes/7
[sprintf]: https://pkg.go.dev/fmt#Sprintf
[format-cheat-sheet]: https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/
[seed]: https://pkg.go.dev/math/rand#Seed
[source]: https://pkg.go.dev/math/rand#Source
[now]: https://pkg.go.dev/time#Now
[unixnano]: https://pkg.go.dev/time#Time.UnixNano
[shuffle]: https://pkg.go.dev/math/rand#Shuffle
[zero-value]: https://yourbasic.org/golang/default-zero-value/
