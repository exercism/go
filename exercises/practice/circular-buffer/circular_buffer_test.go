package circularbuffer

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

// Here is one way you can have a test case verify that the expected
// interfaces are implemented.

var (
	_ io.ByteReader = new(Buffer)
	_ io.ByteWriter = new(Buffer)
)

func (tc testCase) Run(t *testing.T) {
	var calls []string

	b := NewBuffer(tc.capacity)
	calls = append(calls, fmt.Sprintf("b := NewBuffer(%d)", tc.capacity))
	for _, op := range tc.operations {
		var got byte
		var err error
		switch op.name {
		case "ReadByte":
			got, err = b.ReadByte()
			calls = append(calls, "b.ReadByte()")
		case "WriteByte":
			err = b.WriteByte(op.value)
			calls = append(calls, fmt.Sprintf("b.WriteByte(%q)", op.value))
		case "Overwrite":
			b.Overwrite(op.value)
			calls = append(calls, fmt.Sprintf("b.Overwrite(%q)", op.value))
		case "Reset":
			b.Reset()
			calls = append(calls, "b.Reset()")
		}
		if op.name == "ReadByte" || op.name == "WriteByte" {
			if op.wantErr && err != nil {
				continue
			} else if op.wantErr {
				t.Fatalf("calls: %s\nexpected error for the last call but got nil", strings.Join(calls, "; "))
			} else if err != nil {
				t.Fatalf("calls: %s\ngot unexpected error, %v", strings.Join(calls, "; "), err)
			}
		}
		if op.name == "ReadByte" && got != op.value {
			t.Fatalf("calls: %s\ngot %q, want %q", strings.Join(calls, "; "), got, op.value)
		}
	}
}

func TestRun(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) { tc.Run(t) })
	}
}

func BenchmarkRun(b *testing.B) {
	for range b.N {
		for _, tc := range testCases {
			b := NewBuffer(tc.capacity)
			for _, op := range tc.operations {
				switch op.name {
				case "ReadByte":
					b.ReadByte()
				case "WriteByte":
					b.WriteByte(op.value)
				case "Overwrite":
					b.Overwrite(op.value)
				case "Reset":
					b.Reset()
				}
			}
		}
	}
}

func BenchmarkOverwrite(b *testing.B) {
	c := NewBuffer(100)
	b.ResetTimer()
	for range b.N {
		c.Overwrite(0)
	}
	b.SetBytes(int64(b.N))
}

func BenchmarkWriteRead(b *testing.B) {
	c := NewBuffer(100)
	b.ResetTimer()
	for range b.N {
		c.WriteByte(0)
		c.ReadByte()
	}
	b.SetBytes(int64(b.N))
}
