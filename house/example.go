package house

func Embed(relPhrase, nounPhrase string) string {
	return relPhrase + " " + nounPhrase
}

func Verse(subject string, relPhrases []string, nounPhrase string) string {
	return subject + " " + recurse(relPhrases, nounPhrase)
}

func recurse(relPhrases []string, nounPhrase string) string {
	if len(relPhrases) == 0 {
		return nounPhrase
	}
	return Embed(relPhrases[0], recurse(relPhrases[1:], nounPhrase))
}

func Song() string {
	relPhrases := []string{
		"the horse and the hound and the horn\nthat belonged to",
		"the farmer sowing his corn\nthat kept",
		"the rooster that crowed in the morn\nthat woke",
		"the priest all shaven and shorn\nthat married",
		"the man all tattered and torn\nthat kissed",
		"the maiden all forlorn\nthat milked",
		"the cow with the crumpled horn\nthat tossed",
		"the dog\nthat worried",
		"the cat\nthat killed",
		"the rat\nthat ate",
		"the malt\nthat lay in",
	}
	subject := "This is"
	nounPhrase := "the house that Jack built."
	s := subject + " " + nounPhrase
	for c := len(relPhrases) - 1; c >= 0; c-- {
		s += "\n\n" + Verse(subject, relPhrases[c:], nounPhrase)
	}
	return s
}
