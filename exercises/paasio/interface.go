package paasio

import "io"

type ReadCounter interface {
	io.Reader
	// ReadCount returns the total number of bytes successfully read along
	// with the total number of calls to Read().
	ReadCount() (n int64, nops int)
}

type WriteCounter interface {
	io.Writer
	// WriteCount returns the total number of bytes successfully written along
	// with the total number of calls to Write().
	WriteCount() (n int64, nops int)
}

type ReadWriteCounter interface {
	ReadCounter
	WriteCounter
}
