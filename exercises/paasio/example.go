package paasio

import (
	"io"
	"sync"
)

const TestVersion = 1

// NewWriteCounter returns an implementation of WriteCounter.  Calls to
// w.Write() are not guaranteed to be synchronized.
func NewWriteCounter(w io.Writer) WriteCounter {
	return &writeCounter{
		w:       w,
		counter: newCounter(),
	}
}

// NewReadCounter returns an implementation of ReadCounter.  Calls to
// r.Read() are not guaranteed to be synchronized.
func NewReadCounter(r io.Reader) ReadCounter {
	return &readCounter{
		r:       r,
		counter: newCounter(),
	}
}

/*
// NewReadWriteCounter returns an implementation of ReadWriteCounter.  Calls to
// Calls to rw.Write() and rw.Read() are not guaranteed to be synchronized.
func NewReadWriteCounter(rw io.ReadWriter) (ReadWriteCounter, error) {
	return newReadWriteCounter(rw)
}
*/

type readCounter struct {
	r io.Reader
	counter
}

func (rc *readCounter) Read(p []byte) (int, error) {
	m, err := rc.r.Read(p)
	rc.addBytes(m)
	return m, err
}

func (rc *readCounter) ReadCount() (n int64, nops int) {
	return rc.count()
}

type writeCounter struct {
	w io.Writer
	counter
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	m, err := wc.w.Write(p)
	wc.addBytes(m)
	return m, err
}

func (wc *writeCounter) WriteCount() (n int64, nops int) {
	return wc.count()
}

type counter struct {
	bytes int64
	ops   int
	mutex *sync.Mutex
}

func newCounter() counter {
	return counter{mutex: new(sync.Mutex)}
}

func (c *counter) addBytes(n int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.bytes += int64(n)
	c.ops += 1
}

func (c *counter) count() (int64, int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.bytes, c.ops
}

/*
type rwCounter struct {
	*writeCounter
	*readCounter
}

func newIOCounter(w io.Writer, r io.Reader) (*rwCounter, error) {
	rc, err := newReadCounter(r)
	if err != nil {
		return nil, err
	}
	wc, err := newWriteCounter(w)
	if err != nil {
		return nil, err
	}
	wt := &rwCounter{wc, rc}
	return wt, nil
}

func newReadWriteCounter(rw io.ReadWriter) (*rwCounter, error) {
	return newIOCounter(rw, rw)
}
*/
