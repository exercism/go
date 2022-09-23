//nolint:gosec // In the context of this exercise, it is fine to use math.Rand instead of crypto.Rand.
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
		time.Sleep((time.Duration(rand.Intn(10) + 1)) * time.Millisecond)
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
	const tests = 200
	const bucketSize float64 = 0.6
	var got float64
	foundDifferent := false
	var last float64
	numBuckets := int(12.0 / bucketSize)
	buckets := make([]int, numBuckets)
	for i := 0; i < tests; i++ {
		got = GenerateWandEnergy()
		if got < 0.0 || got >= 12.0 {
			t.Fatalf("GenerateWandEnergy() out of range: %f", got)
		}
		if i > 0 && got != last {
			foundDifferent = true
		}
		buckets[int(got/bucketSize)]++
		last = got
	}
	if !foundDifferent {
		t.Fatalf("GenerateWandEnergy() always generates the same number: %f", got)
	}

	var low, high float64
	for i, v := range buckets {
		if v == 0 {
			low = float64(i) * bucketSize
			high = float64(i+1) * bucketSize
			break
		}
	}
	if high != 0.0 {
		t.Errorf("GenerateWandEnergy() results are not uniformly distributed. %.2f to %.2f should contain values.", low, high)
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
