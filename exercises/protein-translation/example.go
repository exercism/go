package protein

func FromCodon(codon string) string {
	switch codon {
	case
		"AUG":
		return "Methionine"
	case
		"UUU",
		"UUC":
		return "Phenylalanine"
	case
		"UUA",
		"UUG":
		return "Leucine"
	case
		"UCU",
		"UCC",
		"UCA",
		"UCG":
		return "Serine"
	case
		"UAU",
		"UAC":
		return "Tyrosine"
	case
		"UGU",
		"UGC":
		return "Cysteine"
	case
		"UGG":
		return "Tryptophan"
	case
		"UAA",
		"UAG",
		"UGA":
		return "STOP"
	}
	return "STOP"
}

func FromRNA(s string) []string {
	var res = ""
	var proteins []string
	for index, codon := range s {
		res += string(codon)
		if index > 0 && (index+1)%3 == 0 {
			tempCodon := FromCodon(res)
			if tempCodon == "STOP" {
				break
			}
			proteins = append(proteins, tempCodon)
			res = ""
		}
	}
	return proteins
}
