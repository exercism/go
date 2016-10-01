package circular

// standard library container/ring package and "one slot open" technique
// as described at WP.

import (
	"container/ring"
	"errors"
)

const testVersion = 4

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

// ReadByte removes one byte from the buffer and returns it.
// ReadByte returns an error if the buffer is empty.
func (b *Buffer) ReadByte() (byte, error) {
	if b.start == b.end {
		return 0, errors.New("buffer empty")
	}
	c := b.start.Value.(byte)
	b.start = b.start.Next()
	return c, nil
}

// WriteByte puts byte c in the buffer there is room.
// WriteByte returns an error if the buffer is full.
func (b *Buffer) WriteByte(c byte) error {
	if b.end.Next() == b.start {
		return errors.New("buffer full")
	}
	b.end.Value = c
	b.end = b.end.Next()
	return nil
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

// Reset puts the buffer in an empty state.
func (b *Buffer) Reset() {
	b.start = b.end
}
