package circular

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
// We chose the provided API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

// Define the Buffer type here.

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
