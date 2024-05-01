// Package eliudseggs contains an example solution for eliuds-eggs exercise.
package eliudseggs

// EggCount determines the actual egg count from the displayed value.
func EggCount(displayValue int) int {
	count := 0
	for displayValue > 0 {
		if (displayValue & 1) == 1 {
			count++
		}
		displayValue >>= 1
	}
	return count
}
