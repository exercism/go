package grep

// Source: exercism/problem-specifications
// Commit: e893e0b grep: Apply new "input" policy (#1182)
// Problem Specifications Version: 1.1.0

var fileContentData = []string{
	"                                                     ",
	" iliad.txt                                           ",
	"   ---------------------------------------------     ",
	"   |Achilles sing, O Goddess! Peleus' son;     |     ",
	"   |His wrath pernicious, who ten thousand woes|     ",
	"   |Caused to Achaia's host, sent many a soul  |     ",
	"   |Illustrious into Ades premature,           |     ",
	"   |And Heroes gave (so stood the will of Jove)|     ",
	"   |To dogs and to all ravening fowls a prey,  |     ",
	"   |When fierce dispute had separated once     |     ",
	"   |The noble Chief Achilles from the son      |     ",
	"   |Of Atreus, Agamemnon, King of men.         |     ",
	"   ---------------------------------------------     ",
	"                                                     ",
	" midsummer-night.txt                                 ",
	"   -----------------------------------------------   ",
	"   |I do entreat your grace to pardon me.        |   ",
	"   |I know not by what power I am made bold,     |   ",
	"   |Nor how it may concern my modesty,           |   ",
	"   |In such a presence here to plead my thoughts;|   ",
	"   |But I beseech your grace that I may know     |   ",
	"   |The worst that may befall me in this case,   |   ",
	"   |If I refuse to wed Demetrius.                |   ",
	"   -----------------------------------------------   ",
	"                                                     ",
	" paradise-lost.txt                                   ",
	"   ------------------------------------------------- ",
	"   |Of Mans First Disobedience, and the Fruit      | ",
	"   |Of that Forbidden Tree, whose mortal tast      | ",
	"   |Brought Death into the World, and all our woe, | ",
	"   |With loss of Eden, till one greater Man        | ",
	"   |Restore us, and regain the blissful Seat,      | ",
	"   |Sing Heav'nly Muse, that on the secret top     | ",
	"   |Of Oreb, or of Sinai, didst inspire            | ",
	"   |That Shepherd, who first taught the chosen Seed| ",
	"   ------------------------------------------------- ",
}

var testCases = []struct {
	description string
	pattern     string
	flags       []string
	files       []string
	expected    []string
}{

	{
		description: "One file, one match, no flags",
		pattern:     "Agamemnon",
		flags:       []string{},
		files:       []string{"iliad.txt"},
		expected:    []string{"Of Atreus, Agamemnon, King of men."},
	},
	{
		description: "One file, one match, print line numbers flag",
		pattern:     "Forbidden",
		flags:       []string{"-n"},
		files:       []string{"paradise-lost.txt"},
		expected:    []string{"2:Of that Forbidden Tree, whose mortal tast"},
	},
	{
		description: "One file, one match, case-insensitive flag",
		pattern:     "FORBIDDEN",
		flags:       []string{"-i"},
		files:       []string{"paradise-lost.txt"},
		expected:    []string{"Of that Forbidden Tree, whose mortal tast"},
	},
	{
		description: "One file, one match, print file names flag",
		pattern:     "Forbidden",
		flags:       []string{"-l"},
		files:       []string{"paradise-lost.txt"},
		expected:    []string{"paradise-lost.txt"},
	},
	{
		description: "One file, one match, match entire lines flag",
		pattern:     "With loss of Eden, till one greater Man",
		flags:       []string{"-x"},
		files:       []string{"paradise-lost.txt"},
		expected:    []string{"With loss of Eden, till one greater Man"},
	},
	{
		description: "One file, one match, multiple flags",
		pattern:     "OF ATREUS, Agamemnon, KIng of MEN.",
		flags:       []string{"-n", "-i", "-x"},
		files:       []string{"iliad.txt"},
		expected:    []string{"9:Of Atreus, Agamemnon, King of men."},
	},
	{
		description: "One file, several matches, no flags",
		pattern:     "may",
		flags:       []string{},
		files:       []string{"midsummer-night.txt"},
		expected:    []string{"Nor how it may concern my modesty,", "But I beseech your grace that I may know", "The worst that may befall me in this case,"},
	},
	{
		description: "One file, several matches, print line numbers flag",
		pattern:     "may",
		flags:       []string{"-n"},
		files:       []string{"midsummer-night.txt"},
		expected:    []string{"3:Nor how it may concern my modesty,", "5:But I beseech your grace that I may know", "6:The worst that may befall me in this case,"},
	},
	{
		description: "One file, several matches, match entire lines flag",
		pattern:     "may",
		flags:       []string{"-x"},
		files:       []string{"midsummer-night.txt"},
		expected:    []string{},
	},
	{
		description: "One file, several matches, case-insensitive flag",
		pattern:     "ACHILLES",
		flags:       []string{"-i"},
		files:       []string{"iliad.txt"},
		expected:    []string{"Achilles sing, O Goddess! Peleus' son;", "The noble Chief Achilles from the son"},
	},
	{
		description: "One file, several matches, inverted flag",
		pattern:     "Of",
		flags:       []string{"-v"},
		files:       []string{"paradise-lost.txt"},
		expected:    []string{"Brought Death into the World, and all our woe,", "With loss of Eden, till one greater Man", "Restore us, and regain the blissful Seat,", "Sing Heav'nly Muse, that on the secret top", "That Shepherd, who first taught the chosen Seed"},
	},
	{
		description: "One file, no matches, various flags",
		pattern:     "Gandalf",
		flags:       []string{"-n", "-l", "-x", "-i"},
		files:       []string{"iliad.txt"},
		expected:    []string{},
	},

	{
		description: "Multiple files, one match, no flags",
		pattern:     "Agamemnon",
		flags:       []string{},
		files:       []string{"iliad.txt", "midsummer-night.txt", "paradise-lost.txt"},
		expected:    []string{"iliad.txt:Of Atreus, Agamemnon, King of men."},
	},
	{
		description: "Multiple files, several matches, no flags",
		pattern:     "may",
		flags:       []string{},
		files:       []string{"iliad.txt", "midsummer-night.txt", "paradise-lost.txt"},
		expected:    []string{"midsummer-night.txt:Nor how it may concern my modesty,", "midsummer-night.txt:But I beseech your grace that I may know", "midsummer-night.txt:The worst that may befall me in this case,"},
	},
	{
		description: "Multiple files, several matches, print line numbers flag",
		pattern:     "that",
		flags:       []string{"-n"},
		files:       []string{"iliad.txt", "midsummer-night.txt", "paradise-lost.txt"},
		expected:    []string{"midsummer-night.txt:5:But I beseech your grace that I may know", "midsummer-night.txt:6:The worst that may befall me in this case,", "paradise-lost.txt:2:Of that Forbidden Tree, whose mortal tast", "paradise-lost.txt:6:Sing Heav'nly Muse, that on the secret top"},
	},
	{
		description: "Multiple files, one match, print file names flag",
		pattern:     "who",
		flags:       []string{"-l"},
		files:       []string{"iliad.txt", "midsummer-night.txt", "paradise-lost.txt"},
		expected:    []string{"iliad.txt", "paradise-lost.txt"},
	},
	{
		description: "Multiple files, several matches, case-insensitive flag",
		pattern:     "TO",
		flags:       []string{"-i"},
		files:       []string{"iliad.txt", "midsummer-night.txt", "paradise-lost.txt"},
		expected:    []string{"iliad.txt:Caused to Achaia's host, sent many a soul", "iliad.txt:Illustrious into Ades premature,", "iliad.txt:And Heroes gave (so stood the will of Jove)", "iliad.txt:To dogs and to all ravening fowls a prey,", "midsummer-night.txt:I do entreat your grace to pardon me.", "midsummer-night.txt:In such a presence here to plead my thoughts;", "midsummer-night.txt:If I refuse to wed Demetrius.", "paradise-lost.txt:Brought Death into the World, and all our woe,", "paradise-lost.txt:Restore us, and regain the blissful Seat,", "paradise-lost.txt:Sing Heav'nly Muse, that on the secret top"},
	},
	{
		description: "Multiple files, several matches, inverted flag",
		pattern:     "a",
		flags:       []string{"-v"},
		files:       []string{"iliad.txt", "midsummer-night.txt", "paradise-lost.txt"},
		expected:    []string{"iliad.txt:Achilles sing, O Goddess! Peleus' son;", "iliad.txt:The noble Chief Achilles from the son", "midsummer-night.txt:If I refuse to wed Demetrius."},
	},
	{
		description: "Multiple files, one match, match entire lines flag",
		pattern:     "But I beseech your grace that I may know",
		flags:       []string{"-x"},
		files:       []string{"iliad.txt", "midsummer-night.txt", "paradise-lost.txt"},
		expected:    []string{"midsummer-night.txt:But I beseech your grace that I may know"},
	},
	{
		description: "Multiple files, one match, multiple flags",
		pattern:     "WITH LOSS OF EDEN, TILL ONE GREATER MAN",
		flags:       []string{"-n", "-i", "-x"},
		files:       []string{"iliad.txt", "midsummer-night.txt", "paradise-lost.txt"},
		expected:    []string{"paradise-lost.txt:4:With loss of Eden, till one greater Man"},
	},
	{
		description: "Multiple files, no matches, various flags",
		pattern:     "Frodo",
		flags:       []string{"-n", "-l", "-x", "-i"},
		files:       []string{"iliad.txt", "midsummer-night.txt", "paradise-lost.txt"},
		expected:    []string{},
	},
}
