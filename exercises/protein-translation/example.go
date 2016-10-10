package protein

const testVersion = 1

func proteinLookUp(codon string) string {
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

func toProtein(s string) string {
	listCodon := []rune(s)
	var res = ""
	var tempCodon = ""
	var proteinName = ""
	for index, codon := range listCodon {
		res += string(codon)
		if index > 0 && (index+1)%3 == 0 {
			tempCodon = proteinLookUp(res)
			if tempCodon == "STOP" {
				break
			}
			if proteinName == "" {
				proteinName += tempCodon
			} else {
				proteinName += " " + tempCodon
			}
			res = ""
		}
	}
	return proteinName
}
