package paasio

import (
	"bytes"
	"crypto/rand"
	"io"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestMultiThreaded(t *testing.T) {
	mincpu := 2
	minproc := 2
	ncpu := runtime.NumCPU()
	if ncpu < mincpu {
		t.Fatalf("at least %d cpu cores are required", mincpu)
	}
	nproc := runtime.GOMAXPROCS(0)
	if nproc < minproc {
		t.Errorf("at least %d threads are required; rerun the tests", minproc)
		t.Errorf("")
		t.Errorf("\tgo test -cpu %d ...", minproc)
	}
}

// this test could be improved to test that error conditions are preserved.
func testWrite(t *testing.T, writer func(io.Writer) WriteCounter) {
	for i, test := range []struct {
		writes []string
	}{
		{nil},
		{[]string{""}},
		{[]string{"I", " ", "never met ", "", "a gohper"}},
	} {
		var buf bytes.Buffer
		buft := writer(&buf)
		for _, s := range test.writes {
			n, err := buft.Write([]byte(s))
			if err != nil {
				t.Errorf("test %d: Write(%q) unexpected error: %v", i, s, err)
				continue
			}
			if n != len(s) {
				t.Errorf("test %d: Write(%q) unexpected number of bytes written: %v", i, s, n)
				continue
			}
		}
		out := buf.String()
		if out != strings.Join(test.writes, "") {
			t.Errorf("test %d: unexpected content in underlying writer: %q", i, out)
		}
	}
}

func TestWriteWriter(t *testing.T) {
	testWrite(t, NewWriteCounter)
}

func TestWriteReadWriter(t *testing.T) {
	testWrite(t, func(w io.Writer) WriteCounter {
		var r nopReader
		return NewReadWriteCounter(readWriter{r, w})
	})
}

// this test could be improved to test exact number of operations as well as
// ensure that error conditions are preserved.
func testRead(t *testing.T, reader func(io.Reader) ReadCounter) {
	chunkLen := 10 << 20 // 10MB
	orig := make([]byte, 10<<20)
	_, err := rand.Read(orig)
	if err != nil {
		t.Fatalf("error reading random data")
	}
	buf := bytes.NewBuffer(orig)
	rc := reader(buf)
	var obuf bytes.Buffer
	ncopy, err := io.Copy(&obuf, rc)
	if err != nil {
		t.Fatalf("error reading: %v", err)
	}
	if ncopy != int64(chunkLen) {
		t.Fatalf("copied %d bytes instead of %d", ncopy, chunkLen)
	}
	if string(orig) != obuf.String() {
		t.Fatalf("unexpected output from Read()")
	}
	n, nops := rc.ReadCount()
	if n != int64(chunkLen) {
		t.Fatalf("reported %d bytes read instead of %d", n, chunkLen)
	}
	if nops < 2 {
		t.Fatalf("unexpected number of reads: %v", nops)
	}
}

func TestReadReader(t *testing.T) {
	testRead(t, NewReadCounter)
}

func TestReadReadWriter(t *testing.T) {
	testRead(t, func(r io.Reader) ReadCounter {
		var w nopWriter
		return NewReadWriteCounter(readWriter{r, w})
	})
}

func testReadTotal(t *testing.T, rc ReadCounter) {
	numGo := 8000
	numBytes := 50
	totalBytes := int64(numGo) * int64(numBytes)
	p := make([]byte, numBytes)

	t.Logf("Calling Read() for %d*%d=%d bytes", numGo, numBytes, totalBytes)
	wg := new(sync.WaitGroup)
	wg.Add(numGo)
	start := make(chan struct{})
	for i := 0; i < numGo; i++ {
		go func() {
			<-start
			rc.Read(p)
			wg.Done()
		}()
	}
	close(start)

	wg.Wait()
	n, nops := rc.ReadCount()
	if n != totalBytes {
		t.Errorf("expected %d bytes read; %d bytes reported", totalBytes, n)
	}
	if nops != numGo {
		t.Errorf("expected %d read operations; %d operations reported", numGo, nops)
	}
}

func TestReadTotalReader(t *testing.T) {
	var r nopReader
	testReadTotal(t, NewReadCounter(r))
}

func TestReadTotalReadWriter(t *testing.T) {
	var rw nopReadWriter
	testReadTotal(t, NewReadWriteCounter(rw))
}

func testWriteTotal(t *testing.T, wt WriteCounter) {
	numGo := 8000
	numBytes := 50
	totalBytes := int64(numGo) * int64(numBytes)
	p := make([]byte, numBytes)

	t.Logf("Calling Write() with %d*%d=%d bytes", numGo, numBytes, totalBytes)
	wg := new(sync.WaitGroup)
	wg.Add(numGo)
	start := make(chan struct{})
	for i := 0; i < numGo; i++ {
		go func() {
			<-start
			wt.Write(p)
			wg.Done()
		}()
	}
	close(start)

	wg.Wait()
	n, nops := wt.WriteCount()
	if n != totalBytes {
		t.Errorf("expected %d bytes written; %d bytes reported", totalBytes, n)
	}
	if nops != numGo {
		t.Errorf("expected %d write operations; %d operations reported", numGo, nops)
	}
}

func TestWriteTotalWriter(t *testing.T) {
	var w nopWriter
	testWriteTotal(t, NewWriteCounter(w))
}

func TestWriteTotalReadWriter(t *testing.T) {
	var rw nopReadWriter
	testWriteTotal(t, NewReadWriteCounter(rw))
}

func TestReadCountConsistencyReader(t *testing.T) {
	var r nopReader
	testReadCountConsistency(t, NewReadCounter(r))
}

func TestReadCountConsistencyReadWriter(t *testing.T) {
	var rw nopReadWriter
	testReadCountConsistency(t, NewReadWriteCounter(rw))
}

func testReadCountConsistency(t *testing.T, rc ReadCounter) {
	const numGo = 4000
	const numBytes = 50
	p := make([]byte, numBytes)

	wg := new(sync.WaitGroup)
	wg.Add(2 * numGo)
	start := make(chan struct{})
	for i := 0; i < numGo; i++ {
		go func() {
			<-start
			rc.Read(p)
			wg.Done()
		}()
		go func() {
			<-start
			n, nops := rc.ReadCount()
			expectedOps := n / numBytes
			if int64(nops) != expectedOps {
				t.Errorf("expected %d ops@%d bytes read; %d ops reported", expectedOps, n, nops)
			}
			wg.Done()
		}()
	}
	close(start)
	wg.Wait()
}

func TestWriteCountConsistencyWriter(t *testing.T) {
	var w nopWriter
	testWriteCountConsistency(t, NewWriteCounter(w))
}

func TestWriteCountConsistencyReadWriter(t *testing.T) {
	var rw nopReadWriter
	testWriteCountConsistency(t, NewReadWriteCounter(rw))
}

func testWriteCountConsistency(t *testing.T, wc WriteCounter) {
	const numGo = 4000
	const numBytes = 50
	p := make([]byte, numBytes)

	wg := new(sync.WaitGroup)
	wg.Add(2 * numGo)
	start := make(chan struct{})
	for i := 0; i < numGo; i++ {
		go func() {
			<-start
			wc.Write(p)
			wg.Done()
		}()
		go func() {
			<-start
			n, nops := wc.WriteCount()
			expectedOps := n / numBytes
			if int64(nops) != n/numBytes {
				t.Errorf("expected %d nops@%d bytes written; %d ops reported", expectedOps, n, nops)
			}
			wg.Done()
		}()
	}
	close(start)
	wg.Wait()
}

type nopWriter struct{ error }

func (w nopWriter) Write(p []byte) (int, error) {
	time.Sleep(1)
	if w.error != nil {
		return 0, w.error
	}
	return len(p), nil
}

type nopReader struct{ error }

func (r nopReader) Read(p []byte) (int, error) {
	time.Sleep(1)
	if r.error != nil {
		return 0, r.error
	}
	return len(p), nil
}

type nopReadWriter struct {
	nopReader
	nopWriter
}

type readWriter struct {
	io.Reader
	io.Writer
}
