package house

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: b820099 Allow prettier to format more files (#1966)

type testCase struct {
	description string
	verse       int
	expected    string
}

var testCases = []testCase{
	{
		description: "verse one - the house that jack built",
		verse:       1,
		expected:    "This is the house that Jack built.",
	},
	{
		description: "verse two - the malt that lay",
		verse:       2,
		expected:    "This is the malt\nthat lay in the house that Jack built.",
	},
	{
		description: "verse three - the rat that ate",
		verse:       3,
		expected:    "This is the rat\nthat ate the malt\nthat lay in the house that Jack built.",
	},
	{
		description: "verse four - the cat that killed",
		verse:       4,
		expected:    "This is the cat\nthat killed the rat\nthat ate the malt\nthat lay in the house that Jack built.",
	},
	{
		description: "verse five - the dog that worried",
		verse:       5,
		expected:    "This is the dog\nthat worried the cat\nthat killed the rat\nthat ate the malt\nthat lay in the house that Jack built.",
	},
	{
		description: "verse six - the cow with the crumpled horn",
		verse:       6,
		expected:    "This is the cow with the crumpled horn\nthat tossed the dog\nthat worried the cat\nthat killed the rat\nthat ate the malt\nthat lay in the house that Jack built.",
	},
	{
		description: "verse seven - the maiden all forlorn",
		verse:       7,
		expected:    "This is the maiden all forlorn\nthat milked the cow with the crumpled horn\nthat tossed the dog\nthat worried the cat\nthat killed the rat\nthat ate the malt\nthat lay in the house that Jack built.",
	},
	{
		description: "verse eight - the man all tattered and torn",
		verse:       8,
		expected:    "This is the man all tattered and torn\nthat kissed the maiden all forlorn\nthat milked the cow with the crumpled horn\nthat tossed the dog\nthat worried the cat\nthat killed the rat\nthat ate the malt\nthat lay in the house that Jack built.",
	},
	{
		description: "verse nine - the priest all shaven and shorn",
		verse:       9,
		expected:    "This is the priest all shaven and shorn\nthat married the man all tattered and torn\nthat kissed the maiden all forlorn\nthat milked the cow with the crumpled horn\nthat tossed the dog\nthat worried the cat\nthat killed the rat\nthat ate the malt\nthat lay in the house that Jack built.",
	},
	{
		description: "verse 10 - the rooster that crowed in the morn",
		verse:       10,
		expected:    "This is the rooster that crowed in the morn\nthat woke the priest all shaven and shorn\nthat married the man all tattered and torn\nthat kissed the maiden all forlorn\nthat milked the cow with the crumpled horn\nthat tossed the dog\nthat worried the cat\nthat killed the rat\nthat ate the malt\nthat lay in the house that Jack built.",
	},
	{
		description: "verse 11 - the farmer sowing his corn",
		verse:       11,
		expected:    "This is the farmer sowing his corn\nthat kept the rooster that crowed in the morn\nthat woke the priest all shaven and shorn\nthat married the man all tattered and torn\nthat kissed the maiden all forlorn\nthat milked the cow with the crumpled horn\nthat tossed the dog\nthat worried the cat\nthat killed the rat\nthat ate the malt\nthat lay in the house that Jack built.",
	},
	{
		description: "verse 12 - the horse and the hound and the horn",
		verse:       12,
		expected:    "This is the horse and the hound and the horn\nthat belonged to the farmer sowing his corn\nthat kept the rooster that crowed in the morn\nthat woke the priest all shaven and shorn\nthat married the man all tattered and torn\nthat kissed the maiden all forlorn\nthat milked the cow with the crumpled horn\nthat tossed the dog\nthat worried the cat\nthat killed the rat\nthat ate the malt\nthat lay in the house that Jack built.",
	},
}

var song = []string{
	"This is the house that Jack built.",
	"This is the malt\nthat lay in the house that Jack built.",
	"This is the rat\nthat ate the malt\nthat lay in the house that Jack built.",
	"This is the cat\nthat killed the rat\nthat ate the malt\nthat lay in the house that Jack built.",
	"This is the dog\nthat worried the cat\nthat killed the rat\nthat ate the malt\nthat lay in the house that Jack built.",
	"This is the cow with the crumpled horn\nthat tossed the dog\nthat worried the cat\nthat killed the rat\nthat ate the malt\nthat lay in the house that Jack built.",
	"This is the maiden all forlorn\nthat milked the cow with the crumpled horn\nthat tossed the dog\nthat worried the cat\nthat killed the rat\nthat ate the malt\nthat lay in the house that Jack built.",
	"This is the man all tattered and torn\nthat kissed the maiden all forlorn\nthat milked the cow with the crumpled horn\nthat tossed the dog\nthat worried the cat\nthat killed the rat\nthat ate the malt\nthat lay in the house that Jack built.",
	"This is the priest all shaven and shorn\nthat married the man all tattered and torn\nthat kissed the maiden all forlorn\nthat milked the cow with the crumpled horn\nthat tossed the dog\nthat worried the cat\nthat killed the rat\nthat ate the malt\nthat lay in the house that Jack built.",
	"This is the rooster that crowed in the morn\nthat woke the priest all shaven and shorn\nthat married the man all tattered and torn\nthat kissed the maiden all forlorn\nthat milked the cow with the crumpled horn\nthat tossed the dog\nthat worried the cat\nthat killed the rat\nthat ate the malt\nthat lay in the house that Jack built.",
	"This is the farmer sowing his corn\nthat kept the rooster that crowed in the morn\nthat woke the priest all shaven and shorn\nthat married the man all tattered and torn\nthat kissed the maiden all forlorn\nthat milked the cow with the crumpled horn\nthat tossed the dog\nthat worried the cat\nthat killed the rat\nthat ate the malt\nthat lay in the house that Jack built.",
	"This is the horse and the hound and the horn\nthat belonged to the farmer sowing his corn\nthat kept the rooster that crowed in the morn\nthat woke the priest all shaven and shorn\nthat married the man all tattered and torn\nthat kissed the maiden all forlorn\nthat milked the cow with the crumpled horn\nthat tossed the dog\nthat worried the cat\nthat killed the rat\nthat ate the malt\nthat lay in the house that Jack built.",
}
