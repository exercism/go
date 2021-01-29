package raindrops

// Source: exercism/problem-specifications
// Commit: 99de15d raindrops: apply "input" policy
// Problem Specifications Version: 1.1.0

var tests = []struct {
	input    int
	expected string
}{
	{1, "1"},
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
	{6, "Pling"},
	{8, "8"},
	{9, "Pling"},
	{10, "Plang"},
	{14, "Plong"},
	{15, "PlingPlang"},
	{21, "PlingPlong"},
	{25, "Plang"},
	{27, "Pling"},
	{35, "PlangPlong"},
	{49, "Plong"},
	{52, "52"},
	{105, "PlingPlangPlong"},
	{3125, "Plang"},
}
