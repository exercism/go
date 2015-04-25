package counter

import "unicode"

// Incorrect implementation: wrongly counts lines.
type Impl1 struct {
	lines, characters, letters int
}

func (c *Impl1) AddString(s string) {
	for _, char := range s {
		if char == '\n' {
			c.lines++
		} else if unicode.IsLetter(char) {
			c.letters++
		}
		c.characters++
	}
}

func (c Impl1) Lines() int      { return c.lines }
func (c Impl1) Letters() int    { return c.letters }
func (c Impl1) Characters() int { return c.characters }
