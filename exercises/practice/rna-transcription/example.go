package strand

import "strings"

func ToRNA(dna string) string {
	dna = strings.ReplaceAll(dna, "A", "u")
	dna = strings.ReplaceAll(dna, "T", "a")
	dna = strings.ReplaceAll(dna, "C", "g")
	dna = strings.ReplaceAll(dna, "G", "c")
	return strings.ToUpper(dna)
}
