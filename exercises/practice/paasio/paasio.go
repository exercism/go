package paasio

import "io"

// Define readCounter and writeCounter types here.

func NewWriteCounter(writer io.Writer) WriteCounter {
	panic("Please implement the NewWriterCounter function")
}

func NewReadCounter(reader io.Reader) ReadCounter {
	panic("Please implement the NewReadCounter function")
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	panic("Please implement the NewReadWriteCounter function")
}

func (rc *readCounter) Read(p []byte) (int, error) {
	panic("Please implement the Read function")
}

func (rc *readCounter) ReadCount() (int64, int) {
	panic("Please implement the ReadCount function")
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	panic("Please implement the Write function")
}

func (wc *writeCounter) WriteCount() (int64, int) {
	panic("Please implement the WriteCount function")
}
