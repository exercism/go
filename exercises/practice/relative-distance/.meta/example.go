package relativedistance

func DegreeOfSeparation(familyTree map[string][]string, personA, personB string) (int, bool) {
	relatives := make(map[string][]string)
	for a, bs := range familyTree {
		for _, b := range bs {
			relatives[a] = append(relatives[a], b)
			relatives[b] = append(relatives[b], a)
			for _, c := range bs {
				if b != c {
					relatives[b] = append(relatives[b], c)
				}
			}
		}
	}

	if _, ok := relatives[personA]; !ok {
		return 0, false
	}
	if _, ok := relatives[personB]; !ok {
		return 0, false
	}

	type element struct {
		distance int
		person   string
	}
	todo := []element{{0, personA}}
	seen := make(map[string]bool)
	for len(todo) > 0 {
		item := todo[0]
		todo = todo[1:]
		if item.person == personB {
			return item.distance, true
		}
		for _, relative := range relatives[item.person] {
			if !seen[relative] {
				seen[relative] = true
				todo = append(todo, element{item.distance + 1, relative})
			}
		}
	}
	return 0, false
}
