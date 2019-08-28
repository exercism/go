package towerofhanoi

// This is an example implementation.

// Solve returns list of moves for solving the Tower of Hanoi
// puzzle. 'disks' is the amount of disks. 'from' is source
// rod, 'aux' is the helper rod, 'to' is the destination rod.
// If Solve is called as Solve(1, "a", "b", "c"), it should
// return []string{"a -> c"}.
func Solve(disks int, from, aux, to string) []string {

	if disks < 1 {
		return nil // ugh
	}

	if disks == 1 {
		// single disk on rod FROM
		return []string{from + " -> " + to}
	}

	// pin largest disk on rod FROM,
	// then move all other disks to rod AUX
	// (using rod TO as auxiliar)
	m1 := Solve(disks-1, from, to, aux)

	// move largest disk to rod TO
	m2 := append(m1, from+" -> "+to)

	// move all remaining disks from rod AUX to rod TO
	// (using rod FROM as auxiliar)
	return append(m2, Solve(disks-1, aux, from, to)...)
}
