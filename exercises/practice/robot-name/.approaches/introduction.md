# Introduction

There are at least a couple ways to solve Robot Name in Go.

One way is to generate the names randomly, add them to a collection of used names, and then check against that collection to prevent duplicate names.
However, when randomly generating all possible names, the chance for collision increases as the amount of used names increases,
so it takes longer and longer to randomly find an unused name, thus risking the tests of timing out.
Another approach is to generate all possible names sequentially, shuffle them in a collection, and then take from the shuffled collection sequentially.

## General guidance

The most challenging objective is to give out the names in random order, but to do so in a relatively quick amount of time.

## Approach: Generate sequentially and shuffle

```go
// Package robotname is a small library for generating robot names.
package robotname

import (
	"fmt"
	"math/rand"
	"time"
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

	rand.Seed(time.Now().UnixNano())
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

For more information, check the [shuffle approach][approach-shuffle].

[approach-shuffle]: https://exercism.org/tracks/go/exercises/robot-name/approaches/shuffle
