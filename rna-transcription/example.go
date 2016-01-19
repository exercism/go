package strand

import "strings"

const TestVersion = 2

func ToRNA(dna string) string {
	dna = strings.Replace(dna, "A", "u", -1)
	dna = strings.Replace(dna, "T", "a", -1)
	dna = strings.Replace(dna, "C", "g", -1)
	dna = strings.Replace(dna, "G", "c", -1)
	return strings.ToUpper(dna)
}
