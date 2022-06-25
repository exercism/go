package chance

import (
	"math/rand"
	"time"
)

// SeedWithTime seeds math/rand with the current computer time
func SeedWithTime() {
	rand.Seed(time.Now().UnixNano())
}

// RollADie returns a random int between 1 and 20
func RollADie() int {
	return 1 + rand.Intn(20)
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0
func GenerateWandEnergy() float64 {
	return 12.0 * rand.Float64()
}

// ShuffleAnimals returns a slice with eight animal strings in random order
func ShuffleAnimals() []string {
	result := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	return result
}
