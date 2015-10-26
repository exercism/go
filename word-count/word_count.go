// +build !example

package wordcount

const TestVersion = 1

// Use this return type.
type Frequency map[string]int

// Just implement the function.
func WordCount(phrase string) Frequency
