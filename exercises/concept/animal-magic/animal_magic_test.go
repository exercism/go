package chance

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestSeedWithTime(t *testing.T) {
	const tests = 100
	var last int64
	for i := 0; i < tests; i++ {
		SeedWithTime()
		got := rand.Int63()
		if i > 0 && got != last {
			return
		}
		last = got
		time.Sleep(time.Duration(rand.Intn(10)) * time.Microsecond)
	}
	t.Errorf("SeedWithTime always sets the same seed")
}

func TestRollADie(t *testing.T) {
	const tests = 100
	var got int
	foundDifferent := false
	var last int
	for i := 0; i < tests; i++ {
		got = RollADie()
		if got < 1 || got > 20 {
			t.Fatalf("RollADie() out of range: %d", got)
		}
		if i > 0 && got != last {
			foundDifferent = true
		}
		last = got
	}
	if !foundDifferent {
		t.Errorf("RollADie() always generates the same number: %d", got)
	}
}

func TestWandEnergy(t *testing.T) {
	const tests = 100
	var got float64
	foundDifferent := false
	var last float64
	for i := 0; i < tests; i++ {
		got = GenerateWandEnergy()
		if got < 0.0 || got >= 12.0 {
			t.Fatalf("GenerateWandEnergy() out of range: %f", got)
		}
		if i > 0 && got != last {
			foundDifferent = true
		}
		last = got
	}
	if !foundDifferent {
		t.Errorf("GenerateWandEnergy() always generates the same number: %f", got)
	}
}

func TestShuffleAnimals(t *testing.T) {
	const tests = 100
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	foundDifferent := false
	var last []string
	var got []string
	for i := 0; i < tests; i++ {
		got = ShuffleAnimals()
		gotSorted := make([]string, len(got))
		copy(gotSorted, got)
		sort.Strings(gotSorted)
		if !slicesEqual(gotSorted, animals) {
			t.Fatalf("ShuffleAnimals() returns incorrect slice: %v", got)
		}
		if i > 0 && !foundDifferent && !slicesEqual(last, got) {
			foundDifferent = true
		}
		last = got
	}
	if !foundDifferent {
		t.Errorf("ShuffleAnimals() always generates the same slice: %v", got)
	}
}

func slicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}
	size := len(a)
	for i := 0; i < size; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
