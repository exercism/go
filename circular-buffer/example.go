package circular

// standard library container/ring package and "one slot open" technique
// as described at WP.

import "container/ring"

// Buffer implements a circular buffer supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
type Buffer struct {
	start, end *ring.Ring
}

// NewBuffer constructs a new empty circular buffer.
func NewBuffer(size int) *Buffer {
	r := ring.New(size + 1)
	return &Buffer{r, r}
}

// Read removes one byte from the buffer and returns it.
// Read returns ok = false if the buffer is empty.
func (b *Buffer) Read() (c byte, ok bool) {
	if b.start == b.end {
		return 0, false
	}
	c = b.start.Value.(byte)
	b.start = b.start.Next()
	return c, true
}

// Write puts byte c in the buffer and returns true as long as there is room
// in the buffer.  Write returns false if the buffer is full.
func (b *Buffer) Write(c byte) bool {
	if b.end.Next() == b.start {
		return false
	}
	b.end.Value = c
	b.end = b.end.Next()
	return true
}

// Overwrite unconditionally puts byte c in the buffer.  If the buffer was
// already full, c overwrites the oldest byte in the buffer.
func (b *Buffer) Overwrite(c byte) {
	b.end.Value = c
	b.end = b.end.Next()
	if b.end == b.start {
		b.start = b.start.Next()
	}
}

// Clear clears the buffer, leaving it empty.
func (b *Buffer) Clear() {
	b.start = b.end
}
