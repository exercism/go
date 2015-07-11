package counter

// Incorrect implementation: wrongly determines characters.
type Impl2 struct {
	newlines, characters, letters int
	lastChar                      rune
}

func (c *Impl2) AddString(s string) {
	for _, char := range s {
		c.lastChar = char
		if char == '\n' {
			c.newlines++
		} else if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
			c.letters++
		}
		c.characters++
	}
}

func (c Impl2) Lines() int {
	switch {
	case c.characters == 0:
		return 0
	case c.lastChar == '\n':
		return c.newlines
	default:
		return c.newlines + 1
	}
}

func (c Impl2) Letters() int    { return c.letters }
func (c Impl2) Characters() int { return c.characters }
