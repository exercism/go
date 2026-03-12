package kindergartengarden

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: b820099 Allow prettier to format more files (#1966)

type lookup struct {
	child  string
	plants []string
	ok     bool
}

type gardenTest struct {
	description string
	diagram     string
	children    []string
	expectError bool
	lookups     []lookup
}

var testCases = []gardenTest{
	{
		description: "garden with single student",
		diagram:     "\n" + "RC\nGG",
		children:    []string{"Alice"},
		expectError: false,
		lookups: []lookup{
			{
				child:  "Alice",
				plants: []string{"radishes", "clover", "grass", "grass"},
				ok:     true,
			},
		},
	},
	{
		description: "different garden with single student",
		diagram:     "\n" + "VC\nRC",
		children:    []string{"Alice"},
		expectError: false,
		lookups: []lookup{
			{
				child:  "Alice",
				plants: []string{"violets", "clover", "radishes", "clover"},
				ok:     true,
			},
		},
	},
	{
		description: "garden with two students",
		diagram:     "\n" + "VVCG\nVVRC",
		children:    []string{"Alice", "Bob"},
		expectError: false,
		lookups: []lookup{
			{
				child:  "Bob",
				plants: []string{"clover", "grass", "radishes", "clover"},
				ok:     true,
			},
		},
	},
	{
		description: "second student's garden",
		diagram:     "\n" + "VVCCGG\nVVCCGG",
		children:    []string{"Alice", "Bob", "Charlie"},
		expectError: false,
		lookups: []lookup{
			{
				child:  "Bob",
				plants: []string{"clover", "clover", "clover", "clover"},
				ok:     true,
			},
		},
	},
	{
		description: "third student's garden",
		diagram:     "\n" + "VVCCGG\nVVCCGG",
		children:    []string{"Alice", "Bob", "Charlie"},
		expectError: false,
		lookups: []lookup{
			{
				child:  "Charlie",
				plants: []string{"grass", "grass", "grass", "grass"},
				ok:     true,
			},
		},
	},
	{
		description: "for Alice, first student's garden",
		diagram:     "\n" + "VRCGVVRVCGGCCGVRGCVCGCGV\nVRCCCGCRRGVCGCRVVCVGCGCV",
		children:    nil,
		expectError: false,
		lookups: []lookup{
			{
				child:  "Alice",
				plants: []string{"violets", "radishes", "violets", "radishes"},
				ok:     true,
			},
		},
	},
	{
		description: "for Bob, second student's garden",
		diagram:     "\n" + "VRCGVVRVCGGCCGVRGCVCGCGV\nVRCCCGCRRGVCGCRVVCVGCGCV",
		children:    nil,
		expectError: false,
		lookups: []lookup{
			{
				child:  "Bob",
				plants: []string{"clover", "grass", "clover", "clover"},
				ok:     true,
			},
		},
	},
	{
		description: "for Charlie",
		diagram:     "\n" + "VRCGVVRVCGGCCGVRGCVCGCGV\nVRCCCGCRRGVCGCRVVCVGCGCV",
		children:    nil,
		expectError: false,
		lookups: []lookup{
			{
				child:  "Charlie",
				plants: []string{"violets", "violets", "clover", "grass"},
				ok:     true,
			},
		},
	},
	{
		description: "for David",
		diagram:     "\n" + "VRCGVVRVCGGCCGVRGCVCGCGV\nVRCCCGCRRGVCGCRVVCVGCGCV",
		children:    nil,
		expectError: false,
		lookups: []lookup{
			{
				child:  "David",
				plants: []string{"radishes", "violets", "clover", "radishes"},
				ok:     true,
			},
		},
	},
	{
		description: "for Eve",
		diagram:     "\n" + "VRCGVVRVCGGCCGVRGCVCGCGV\nVRCCCGCRRGVCGCRVVCVGCGCV",
		children:    nil,
		expectError: false,
		lookups: []lookup{
			{
				child:  "Eve",
				plants: []string{"clover", "grass", "radishes", "grass"},
				ok:     true,
			},
		},
	},
	{
		description: "for Fred",
		diagram:     "\n" + "VRCGVVRVCGGCCGVRGCVCGCGV\nVRCCCGCRRGVCGCRVVCVGCGCV",
		children:    nil,
		expectError: false,
		lookups: []lookup{
			{
				child:  "Fred",
				plants: []string{"grass", "clover", "violets", "clover"},
				ok:     true,
			},
		},
	},
	{
		description: "for Ginny",
		diagram:     "\n" + "VRCGVVRVCGGCCGVRGCVCGCGV\nVRCCCGCRRGVCGCRVVCVGCGCV",
		children:    nil,
		expectError: false,
		lookups: []lookup{
			{
				child:  "Ginny",
				plants: []string{"clover", "grass", "grass", "clover"},
				ok:     true,
			},
		},
	},
	{
		description: "for Harriet",
		diagram:     "\n" + "VRCGVVRVCGGCCGVRGCVCGCGV\nVRCCCGCRRGVCGCRVVCVGCGCV",
		children:    nil,
		expectError: false,
		lookups: []lookup{
			{
				child:  "Harriet",
				plants: []string{"violets", "radishes", "radishes", "violets"},
				ok:     true,
			},
		},
	},
	{
		description: "for Ileana",
		diagram:     "\n" + "VRCGVVRVCGGCCGVRGCVCGCGV\nVRCCCGCRRGVCGCRVVCVGCGCV",
		children:    nil,
		expectError: false,
		lookups: []lookup{
			{
				child:  "Ileana",
				plants: []string{"grass", "clover", "violets", "clover"},
				ok:     true,
			},
		},
	},
	{
		description: "for Joseph",
		diagram:     "\n" + "VRCGVVRVCGGCCGVRGCVCGCGV\nVRCCCGCRRGVCGCRVVCVGCGCV",
		children:    nil,
		expectError: false,
		lookups: []lookup{
			{
				child:  "Joseph",
				plants: []string{"violets", "clover", "violets", "grass"},
				ok:     true,
			},
		},
	},
	{
		description: "for Kincaid, second to last student's garden",
		diagram:     "\n" + "VRCGVVRVCGGCCGVRGCVCGCGV\nVRCCCGCRRGVCGCRVVCVGCGCV",
		children:    nil,
		expectError: false,
		lookups: []lookup{
			{
				child:  "Kincaid",
				plants: []string{"grass", "clover", "clover", "grass"},
				ok:     true,
			},
		},
	},
	{
		description: "for Larry, last student's garden",
		diagram:     "\n" + "VRCGVVRVCGGCCGVRGCVCGCGV\nVRCCCGCRRGVCGCRVVCVGCGCV",
		children:    nil,
		expectError: false,
		lookups: []lookup{
			{
				child:  "Larry",
				plants: []string{"grass", "violets", "clover", "violets"},
				ok:     true,
			},
		},
	},
}
