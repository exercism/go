package gigasecond

import (
	"math/rand"
	"testing"
	"time"
)

func TestGigasecond(t *testing.T) {
	if Gigasecond != 1e18 {
		t.Fatalf("Gigasecond is incorrectly defined.")
	}
}

func TestAddGigasecond(t *testing.T) {
	for i := 0; i < 100; i++ {
		ns := rand.Int63()
		input := time.Unix(0, ns)
		expect := time.Unix(1e9, ns)
		output := AddGigasecond(input)
		if !output.Equal(expect) {
			t.Fatalf("AddGigasecond(%v): got %v; want %v", input, output, expect)
		}
	}
}

func BenchmarkAddGigasecond(b *testing.B) {
	t := time.Unix(123456789, 123456789)
	for i := 0; i < b.N; i++ {
		AddGigasecond(t)
	}
}
