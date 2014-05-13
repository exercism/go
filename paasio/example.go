package paasio

import (
	"fmt"
	"io"
	"sync"
)

// NewWriteCounter returns an implementation of WriteCounter.  Calls to
// w.Write() are not guaranteed to be synchronized.
func NewWriteCounter(w io.Writer) (WriteCounter, error) {
	return newWriteCounter(w)
}

// NewReadCounter returns an implementation of ReadCounter.  Calls to
// r.Read() are not guaranteed to be synchronized.
func NewReadCounter(r io.Reader) (ReadCounter, error) {
	return newReadCounter(r)
}

// NewReadWriteCounter returns an implementation of ReadWriteCounter.  Calls to
// Calls to rw.Write() and rw.Read() are not guaranteed to be synchronized.
func NewReadWriteCounter(rw io.ReadWriter) (ReadWriteCounter, error) {
	return newReadWriteCounter(rw)
}

type readCounter struct {
	r      io.Reader
	nrd    int64
	nrdops int
	rdl    *sync.Mutex
}

func newReadCounter(r io.Reader) (*readCounter, error) {
	if r == nil {
		return nil, fmt.Errorf("nil reader")
	}
	wt := &readCounter{
		r:   r,
		rdl: new(sync.Mutex),
	}
	return wt, nil
}

func (rc *readCounter) Read(p []byte) (int, error) {
	m, err := rc.r.Read(p)
	rc.rdl.Lock()
	rc.nrd += int64(m)
	rc.nrdops++
	rc.rdl.Unlock()
	return m, err
}

func (rc *readCounter) ReadCount() (n int64, nops int) {
	rc.rdl.Lock()
	n, nops = rc.nrd, rc.nrdops
	rc.rdl.Unlock()
	return n, nops
}

type writeCounter struct {
	w      io.Writer
	nwr    int64
	nwrops int
	wrl    *sync.Mutex
}

func newWriteCounter(w io.Writer) (*writeCounter, error) {
	if w == nil {
		return nil, fmt.Errorf("nil writer")
	}
	wt := &writeCounter{
		w:   w,
		wrl: new(sync.Mutex),
	}
	return wt, nil
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	m, err := wc.w.Write(p)
	wc.wrl.Lock()
	wc.nwr += int64(m)
	wc.nwrops++
	wc.wrl.Unlock()
	return m, err
}

func (wc *writeCounter) WriteCount() (n int64, nops int) {
	wc.wrl.Lock()
	n, nops = wc.nwr, wc.nwrops
	wc.wrl.Unlock()
	return n, nops
}

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
