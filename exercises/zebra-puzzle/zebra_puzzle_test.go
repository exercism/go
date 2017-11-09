package zebra_puzzle

import "testing"

func TestSolve(t *testing.T) {
	const (
		expectedWaterDrinker = "Norwegian"
		expectedZebraOwner   = "Japanese"
	)
	waterDrinker, zebraOwner := Solve()
	if waterDrinker != expectedWaterDrinker {
		t.Fatalf("FAILED: For waterDrinker, expected %q, actual %q", expectedWaterDrinker, waterDrinker)
	}
	if zebraOwner != expectedZebraOwner {
		t.Fatalf("FAILED: For zebraOwner, expected %q, actual %q", expectedZebraOwner, zebraOwner)
	}
}

func BenchmarkScore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solve()
	}
}
