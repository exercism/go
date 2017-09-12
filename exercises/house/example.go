package house

var songLines = []string{
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

// Recursive Solution

func Verse(v int) (verse string) {
	v--
	i := len(songLines) - v
	verse += buildVerse(songLines[i:], "This is ")
	verse += "the house that Jack built."
	return
}

func buildVerse(songLines []string, cur string) string {
	if len(songLines) == 0 {
		return cur
	}
	cur += songLines[0]
	cur += " "
	return buildVerse(songLines[1:], cur)
}

func Song() (song string) {
	for i := 0; i <= len(songLines); i++ {
		song += Verse(i + 1)
		if i < len(songLines) {
			song += "\n\n"
		}
	}
	return
}

// Iterative Solution

// func Verse(v int) (verse string) {
//     v--
//     verse += "This is "
//     verse += strings.Join(songLines[len(songLines)-v:], " ")
//     if v > 0 {
//         verse += " "
//     }
//     verse += "the house that Jack built."
//     return
// }
//
// func Song() (song string) {
//     for i := 0; i <= len(songLines); i++ {
//         song += Verse(i + 1)
//         if i < len(songLines) {
//             song += "\n\n"
//         }
//     }
//     return
// }
