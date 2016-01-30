package paasio

import (
	"io"
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
	rc.updates <- int64(m)
	return m, err
}

func (rc *readCounter) ReadCount() (n int64, nops int) {
	return rc.query()
}

type writeCounter struct {
	w io.Writer
	counter
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	m, err := wc.w.Write(p)
	wc.updates <- int64(m)
	return m, err
}

func (wc *writeCounter) WriteCount() (n int64, nops int) {
	return wc.query()
}

type counts struct {
	bytes int64
	ops   int
}

type counter struct {
	counts
	updates chan int64
	queries chan chan counts
}

func newCounter() counter {
	c := counter{
		updates: make(chan int64),
		queries: make(chan chan counts),
	}

	go func() {
		for {
			select {
			case bytes := <-c.updates:
				c.bytes += bytes
				c.ops += 1
			case dest := <-c.queries:
				dest <- c.counts
			}
		}
	}()

	return c
}

func (c counter) query() (int64, int) {
	ch := make(chan counts)
	c.queries <- ch
	counts := <-ch
	return counts.bytes, counts.ops
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
