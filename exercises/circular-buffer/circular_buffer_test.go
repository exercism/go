package circular

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
//   type Buffer
//   func NewBuffer(size int) *Buffer
//   func (*Buffer) ReadByte() (byte, error)
//   func (*Buffer) WriteByte(c byte) error
//   func (*Buffer) Overwrite(c byte)
//   func (*Buffer) Reset() // put buffer in an empty state
//
// We chose the above API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

import (
	"io"
	"testing"
)

const targetTestVersion = 4

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Errorf("Found testVersion = %v, want %v.", testVersion, targetTestVersion)
	}
}

// Here is one way you can have a test case verify that the expected
// interfaces are implemented.

var _ io.ByteReader = new(Buffer)
var _ io.ByteWriter = new(Buffer)

// testBuffer and methods support the tests, providing log and fail messages.

type testBuffer struct {
	*testing.T
	b *Buffer
}

func nb(size int, t *testing.T) testBuffer {
	t.Logf("NewBuffer(%d)", size)
	return testBuffer{t, NewBuffer(size)}
}

func (tb testBuffer) read(want byte) {
	switch c, err := tb.b.ReadByte(); {
	case err != nil:
		var _ error = err
		tb.Fatalf("ReadByte() failed unexpectedly: %v", err)
	case c != want:
		tb.Fatalf("ReadByte() = %c, want %c.", c, want)
	}
	tb.Logf("ReadByte %c", want)
}

func (tb testBuffer) readFail() {
	c, err := tb.b.ReadByte()
	if err == nil {
		tb.Fatalf("ReadByte() = %c, expected a failure", c)
	}
	var _ error = err
	tb.Log("ReadByte() fails as expected")
}

func (tb testBuffer) write(c byte) {
	if err := tb.b.WriteByte(c); err != nil {
		var _ error = err
		tb.Fatalf("WriteByte(%c) failed unexpectedly: %v", c, err)
	}
	tb.Logf("WriteByte(%c)", c)
}

func (tb testBuffer) writeFail(c byte) {
	err := tb.b.WriteByte(c)
	if err == nil {
		tb.Fatalf("WriteByte(%c) succeeded, expected a failure", c)
	}
	var _ error = err
	tb.Logf("WriteByte(%c) fails as expected", c)
}

func (tb testBuffer) reset() {
	tb.b.Reset()
	tb.Log("Reset()")
}

func (tb testBuffer) overwrite(c byte) {
	tb.b.Overwrite(c)
	tb.Logf("Overwrite(%c)", c)
}

// tests.  separate functions so log will have descriptive test name.

func TestReadEmptyBuffer(t *testing.T) {
	tb := nb(1, t)
	tb.readFail()
}

func TestWriteAndReadOneItem(t *testing.T) {
	tb := nb(1, t)
	tb.write('1')
	tb.read('1')
	tb.readFail()
}

func TestWriteAndReadMultipleItems(t *testing.T) {
	tb := nb(2, t)
	tb.write('1')
	tb.write('2')
	tb.read('1')
	tb.read('2')
	tb.readFail()
}

func TestReset(t *testing.T) {
	tb := nb(3, t)
	tb.write('1')
	tb.write('2')
	tb.write('3')
	tb.reset()
	tb.write('1')
	tb.write('3')
	tb.read('1')
	tb.write('4')
	tb.read('3')
}

func TestAlternateWriteAndRead(t *testing.T) {
	tb := nb(2, t)
	tb.write('1')
	tb.read('1')
	tb.write('2')
	tb.read('2')
}

func TestReadOldestItem(t *testing.T) {
	tb := nb(3, t)
	tb.write('1')
	tb.write('2')
	tb.read('1')
	tb.write('3')
	tb.read('2')
	tb.read('3')
}

func TestWriteFullBuffer(t *testing.T) {
	tb := nb(2, t)
	tb.write('1')
	tb.write('2')
	tb.writeFail('A')
}

func TestOverwriteFull(t *testing.T) {
	tb := nb(2, t)
	tb.write('1')
	tb.write('2')
	tb.overwrite('A')
	tb.read('2')
	tb.read('A')
	tb.readFail()
}

func TestOverwriteNonFull(t *testing.T) {
	tb := nb(2, t)
	tb.write('1')
	tb.overwrite('2')
	tb.read('1')
	tb.read('2')
	tb.readFail()
}

func TestAlternateReadAndOverwrite(t *testing.T) {
	tb := nb(5, t)
	tb.write('1')
	tb.write('2')
	tb.write('3')
	tb.read('1')
	tb.read('2')
	tb.write('4')
	tb.read('3')
	tb.write('5')
	tb.write('6')
	tb.write('7')
	tb.write('8')
	tb.overwrite('A')
	tb.overwrite('B')
	tb.read('6')
	tb.read('7')
	tb.read('8')
	tb.read('A')
	tb.read('B')
	tb.readFail()
}

func BenchmarkOverwrite(b *testing.B) {
	c := NewBuffer(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Overwrite(0)
	}
	b.SetBytes(int64(b.N))
}

func BenchmarkWriteRead(b *testing.B) {
	c := NewBuffer(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.WriteByte(0)
		c.ReadByte()
	}
	b.SetBytes(int64(b.N))
}
