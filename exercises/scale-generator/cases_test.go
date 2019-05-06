package scale

// Source: exercism/problem-specifications
// Commit: 65cc51b Update scale-generator version to 2.0.0
// Problem Specifications Version: 2.0.0

type scaleTest struct {
	description string
	tonic       string
	interval    string
	expected    []string
}

var scaleTestCases = []scaleTest{
	{
		description: "Chromatic scale with sharps",
		tonic:       "C",
		interval:    "",
		expected:    []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"},
	},
	{
		description: "Chromatic scale with flats",
		tonic:       "F",
		interval:    "",
		expected:    []string{"F", "Gb", "G", "Ab", "A", "Bb", "B", "C", "Db", "D", "Eb", "E"},
	},
	{
		description: "Simple major scale",
		tonic:       "C",
		interval:    "MMmMMMm",
		expected:    []string{"C", "D", "E", "F", "G", "A", "B"},
	},
	{
		description: "Major scale with sharps",
		tonic:       "G",
		interval:    "MMmMMMm",
		expected:    []string{"G", "A", "B", "C", "D", "E", "F#"},
	},
	{
		description: "Major scale with flats",
		tonic:       "F",
		interval:    "MMmMMMm",
		expected:    []string{"F", "G", "A", "Bb", "C", "D", "E"},
	},
	{
		description: "Minor scale with sharps",
		tonic:       "f#",
		interval:    "MmMMmMM",
		expected:    []string{"F#", "G#", "A", "B", "C#", "D", "E"},
	},
	{
		description: "Minor scale with flats",
		tonic:       "bb",
		interval:    "MmMMmMM",
		expected:    []string{"Bb", "C", "Db", "Eb", "F", "Gb", "Ab"},
	},
	{
		description: "Dorian mode",
		tonic:       "d",
		interval:    "MmMMMmM",
		expected:    []string{"D", "E", "F", "G", "A", "B", "C"},
	},
	{
		description: "Mixolydian mode",
		tonic:       "Eb",
		interval:    "MMmMMmM",
		expected:    []string{"Eb", "F", "G", "Ab", "Bb", "C", "Db"},
	},
	{
		description: "Lydian mode",
		tonic:       "a",
		interval:    "MMMmMMm",
		expected:    []string{"A", "B", "C#", "D#", "E", "F#", "G#"},
	},
	{
		description: "Phrygian mode",
		tonic:       "e",
		interval:    "mMMMmMM",
		expected:    []string{"E", "F", "G", "A", "B", "C", "D"},
	},
	{
		description: "Locrian mode",
		tonic:       "g",
		interval:    "mMMmMMM",
		expected:    []string{"G", "Ab", "Bb", "C", "Db", "Eb", "F"},
	},
	{
		description: "Harmonic minor",
		tonic:       "d",
		interval:    "MmMMmAm",
		expected:    []string{"D", "E", "F", "G", "A", "Bb", "Db"},
	},
	{
		description: "Octatonic",
		tonic:       "C",
		interval:    "MmMmMmMm",
		expected:    []string{"C", "D", "D#", "F", "F#", "G#", "A", "B"},
	},
	{
		description: "Hexatonic",
		tonic:       "Db",
		interval:    "MMMMMM",
		expected:    []string{"Db", "Eb", "F", "G", "A", "B"},
	},
	{
		description: "Pentatonic",
		tonic:       "A",
		interval:    "MMAMA",
		expected:    []string{"A", "B", "C#", "E", "F#"},
	},
	{
		description: "Enigmatic",
		tonic:       "G",
		interval:    "mAMMMmm",
		expected:    []string{"G", "G#", "B", "C#", "D#", "F", "F#"},
	},
}
