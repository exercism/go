package bowling

// Source: exercism/problem-specifications
// Commit: daf84d6 bowling, transpose: conform array format to rest of file

var scoreTestCases = []struct {
	description   string
	previousRolls []int  // bowling rolls to do before the Score() test
	valid         bool   // true => no error, false => error expected
	score         int    // when .valid == true, the expected score value
	explainText   string // when .valid == false, error explanation text
}{}

var rollTestCases = []struct {
	description   string
	previousRolls []int  // bowling rolls to do before the Roll(roll) test
	valid         bool   // true => no error, false => error expected
	roll          int    // pin count for the test roll
	explainText   string // when .valid == false, error explanation text
}{}
