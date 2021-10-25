package circular

import "container/ring"

type Buffer struct {
	start, end *ring.Ring
}

func NewBuffer(size int) *Buffer {
	panic("Please implement the NewBuffer function")
}

func (b *Buffer) ReadByte() (byte, error) {
	panic("Please implement the ReadByte function")
}

func (b *Buffer) WriteByte(c byte) error {
	panic("Please implement the WriteByte function")
}

func (b *Buffer) Overwrite(c byte) {
	panic("Please implement the Overwrite function")
}

func (b *Buffer) Reset() {
	panic("Please implement the Reset function")
}
