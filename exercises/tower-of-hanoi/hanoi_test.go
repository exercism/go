package towerofhanoi

import (
	"fmt"
	"strings"
	"testing"
)

type board struct {
	a []int
	b []int
	c []int
}

func TestTowerOfHanoi(t *testing.T) {

	for i := 0; i < 8; i++ {
		if err := testDisks(t, i); err != nil {
			t.Errorf("FAIL totalDisks=%d %v", i, err)
			continue
		}
		t.Logf("PASS totalDisks=%d", i)
	}
}

func testDisks(t *testing.T, disks int) error {
	// create board
	b := board{}
	for i := disks; i > 0; i-- {
		b.a = append(b.a, i) // push disks into rod A
	}

	// get moves from solution
	moves := Solve(disks, "a", "b", "c")

	// apply moves to board
	err := apply(t, disks, &b, moves)
	if err != nil {
		return err
	}

	// check board
	if len(b.a) != 0 {
		return fmt.Errorf("rod A has %d disks: %v", len(b.a), b.a)
	}
	if len(b.b) != 0 {
		return fmt.Errorf("rod B has %d disks: %v", len(b.b), b.b)
	}
	if len(b.c) != disks {
		return fmt.Errorf("rod C has %d disks: %v", len(b.c), b.c)
	}

	return nil
}

func apply(t *testing.T, disks int, b *board, moves []string) error {
	for i, m := range moves {
		t.Logf("disks=%d move=%s before: %v", disks, m, *b)
		s := strings.Split(m, "->")
		if len(s) != 2 {
			return fmt.Errorf("invalid move format %d: [%s]", i, m)
		}
		rod1 := strings.TrimSpace(s[0])
		rod2 := strings.TrimSpace(s[1])
		if rod1 == rod2 {
			return fmt.Errorf("invalid null move %d: [%s]", i, m)
		}
		var disk int
		switch rod1 {
		case "a":
			if len(b.a) == 0 {
				return fmt.Errorf("invalid move from empty rod A %d: [%s]", i, m)
			}
			last := len(b.a) - 1
			disk, b.a = b.a[last], b.a[:last]
		case "b":
			if len(b.b) == 0 {
				return fmt.Errorf("invalid move from empty rod B %d: [%s]", i, m)
			}
			last := len(b.b) - 1
			disk, b.b = b.b[last], b.b[:last]
		case "c":
			if len(b.c) == 0 {
				return fmt.Errorf("invalid move from empty rod C %d: [%s]", i, m)
			}
			last := len(b.c) - 1
			disk, b.c = b.c[last], b.c[:last]
		default:
			return fmt.Errorf("invalid source rod %d: [%s]", i, m)
		}
		switch rod2 {
		case "a":
			if len(b.a) > 0 {
				last := len(b.a) - 1
				top := b.a[last]
				if disk > top {
					return fmt.Errorf("invalid move disk=%d is greater than top=%d of rod A %d: [%s]", disk, top, i, m)
				}
			}
			b.a = append(b.a, disk)
		case "b":
			if len(b.b) > 0 {
				last := len(b.b) - 1
				top := b.b[last]
				if disk > top {
					return fmt.Errorf("invalid move disk=%d is greater than top=%d of rod A %d: [%s]", disk, top, i, m)
				}
			}
			b.b = append(b.b, disk)
		case "c":
			if len(b.c) > 0 {
				last := len(b.c) - 1
				top := b.c[last]
				if disk > top {
					return fmt.Errorf("invalid move disk=%d is greater than top=%d of rod A %d: [%s]", disk, top, i, m)
				}
			}
			b.c = append(b.c, disk)
		default:
			return fmt.Errorf("invalid destination rod %d: [%s]", i, m)
		}
		t.Logf("disks=%d move=%s after: %v", disks, m, *b)
	}
	return nil
}
