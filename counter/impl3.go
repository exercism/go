package counter

import "unicode"

// Incorrect implementation: assumes ASCII.
type Impl3 struct {
	newlines, characters, letters int
	lastChar                      rune
}

func (c *Impl3) AddString(s string) {
	for i := 0; i < len(s); i++ {
		char := rune(s[i])
		c.lastChar = char
		if char == '\n' {
			c.newlines++
		} else if unicode.IsLetter(char) {
			c.letters++
		}
		c.characters++
	}
}

func (c Impl3) Lines() int {
	switch {
	case c.characters == 0:
		return 0
	case c.lastChar == '\n':
		return c.newlines
	default:
		return c.newlines + 1
	}
}

func (c Impl3) Letters() int    { return c.letters }
func (c Impl3) Characters() int { return c.characters }
