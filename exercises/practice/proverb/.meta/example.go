// Package proverb deals with outputting the lines of a rhyme.
package proverb

import "fmt"

// Proverb outputs the lines of for want of a nail
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}
	var r []string
	if len(rhyme) > 1 {
		for i := 0; i < len(rhyme)-1; i++ {
			r = append(r, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
		}
	}
	r = append(r, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	return r
}
