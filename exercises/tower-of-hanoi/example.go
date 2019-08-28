package towerofhanoi

// This is an example implementation.

// Solve returns list of moves for solving the Tower of Hanoi
// puzzle. 'disks' is the amount of disks. 'from' is the source
// rod, 'to' is the destination rod, 'via' is a helper rod.
// If Solve is called as Solve(1, "a", "b", "c"), it should
// return []string{"a -> b"}.
func Solve(disks int, from, to, via string) []string {

	if disks < 1 {
		return nil // ugh
	}

	fromTo := from + " -> " + to

	if disks == 1 {
		// single disk on rod FROM
		return []string{fromTo}
	}

	// pin largest disk on rod FROM,
	// then move all other disks to rod VIA
	// (using rod TO as helper)
	m1 := Solve(disks-1, from, via, to)

	// move largest disk to rod TO
	m2 := append(m1, fromTo)

	// move all remaining disks from rod VIA to rod TO
	// (using rod FROM as helper)
	return append(m2, Solve(disks-1, via, to, from)...)
}
