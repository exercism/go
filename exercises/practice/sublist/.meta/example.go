package sublist

import "reflect"

// Sublist checks difference of two lists and
// returns equal, sublist, superlist or unequal according
// to their relation to each other.
func Sublist(l1, l2 []int) Relation {
	switch {
	case reflect.DeepEqual(l1, l2):
		return RelationEqual
	case contains(l1, l2):
		return RelationSuperlist
	case contains(l2, l1):
		return RelationSublist
	default:
		return RelationUnequal
	}
}

func contains(l1, l2 []int) bool {
	if len(l2) == 0 {
		return true
	} else if len(l2) > len(l1) {
		return false
	}
	for i := 0; i <= len(l1)-len(l2); i++ {
		if l1[i] != l2[0] {
			continue
		}
		rest := true
		for j, e := range l2 {
			rest = rest && l1[i+j] == e
		}
		if rest {
			return true
		}
	}
	return false
}
