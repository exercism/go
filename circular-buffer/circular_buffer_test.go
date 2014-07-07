package circular

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
//   type Buffer
//   func NewBuffer(size int) *Buffer
//   func (*Buffer) Read() (c byte, ok bool)
//   func (*Buffer) Write(c byte) (ok bool)
//   func (*Buffer) Overwrite(c byte)
//   func (*Buffer) Clear() // clear entire buffer so that it is empty

import "testing"

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
	switch c, ok := tb.b.Read(); {
	case !ok:
		tb.Fatal("Read() returned ok = false, want true.")
	case c != want:
		tb.Fatalf("Read() = %c, want %c.", c, want)
	}
	tb.Logf("Read %c", want)
}

func (tb testBuffer) readFail() {
	if c, ok := tb.b.Read(); ok {
		tb.Fatalf("Read() = %c, ok want !ok", c)
	}
	tb.Log("Read() fail")
}

func (tb testBuffer) write(c byte) {
	if !tb.b.Write(c) {
		tb.Fatalf("Write(%c) returned !ok, want ok.", c)
	}
	tb.Logf("Write(%c)", c)
}

func (tb testBuffer) writeFail(c byte) {
	if tb.b.Write(c) {
		tb.Fatalf("Write(%c) returned ok, want !ok.", c)
	}
	tb.Logf("Write(%c) fail", c)
}

func (tb testBuffer) clear() {
	tb.b.Clear()
	tb.Log("Clear()")
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

func TestClear(t *testing.T) {
	tb := nb(3, t)
	tb.write('1')
	tb.write('2')
	tb.write('3')
	tb.clear()
	tb.write('1')
	tb.write('2')
	tb.read('1')
	tb.write('3')
	tb.read('2')
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

func TestOverwrite(t *testing.T) {
	tb := nb(2, t)
	tb.write('1')
	tb.write('2')
	tb.overwrite('A')
	tb.read('2')
	tb.read('A')
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
}

func BenchmarkWriteRead(b *testing.B) {
	c := NewBuffer(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Write(0)
		c.Read()
	}
}
