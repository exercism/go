package counter

import "unicode"

// Correct implementation
type Impl4 struct {
	newlines, characters, letters int
	lastChar                      rune
}

func (c *Impl4) AddString(s string) {
	for _, char := range s {
		c.lastChar = char
		if char == '\n' {
			c.newlines++
		} else if unicode.IsLetter(char) {
			c.letters++
		}
		c.characters++
	}
}

func (c Impl4) Lines() int {
	switch {
	case c.characters == 0:
		return 0
	case c.lastChar == '\n':
		return c.newlines
	default:
		return c.newlines + 1
	}
}

func (c Impl4) Letters() int    { return c.letters }
func (c Impl4) Characters() int { return c.characters }
