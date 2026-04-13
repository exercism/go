package gradeschool

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 7a8722a Reorder keys (#1960)

type student struct {
	name  string
	grade int
}

type testCaseAdd struct {
	description string
	students    []student
	expected    []bool
}

var testCasesAdd = []testCaseAdd{
	{
		description: "Add a student",
		students: []student{
			student{"Aimee", 2},
		},
		expected: []bool{true},
	},
	{
		description: "Adding multiple students in the same grade in the roster",
		students: []student{
			student{"Blair", 2},
			student{"James", 2},
			student{"Paul", 2},
		},
		expected: []bool{true, true, true},
	},
	{
		description: "Cannot add student to same grade in the roster more than once",
		students: []student{
			student{"Blair", 2},
			student{"James", 2},
			student{"James", 2},
			student{"Paul", 2},
		},
		expected: []bool{true, true, false, true},
	},
	{
		description: "Adding students in multiple grades",
		students: []student{
			student{"Chelsea", 3},
			student{"Logan", 7},
		},
		expected: []bool{true, true},
	},
	{
		description: "Cannot add same student to multiple grades in the roster",
		students: []student{
			student{"Blair", 2},
			student{"James", 2},
			student{"James", 3},
			student{"Paul", 3},
		},
		expected: []bool{true, true, false, true},
	},
}

type testCaseEnrollment struct {
	description string
	students    []student
	expected    []string
}

var testCasesEnrollment = []testCaseEnrollment{
	{
		description: "Roster is empty when no student is added",
		students:    []student{},
		expected:    []string{},
	},
	{
		description: "Student is added to the roster",
		students: []student{
			student{"Aimee", 2},
		},
		expected: []string{"Aimee"},
	},
	{
		description: "Multiple students in the same grade are added to the roster",
		students: []student{
			student{"Blair", 2},
			student{"James", 2},
			student{"Paul", 2},
		},
		expected: []string{"Blair", "James", "Paul"},
	},
	{
		description: "Student not added to same grade in the roster more than once",
		students: []student{
			student{"Blair", 2},
			student{"James", 2},
			student{"James", 2},
			student{"Paul", 2},
		},
		expected: []string{"Blair", "James", "Paul"},
	},
	{
		description: "Students in multiple grades are added to the roster",
		students: []student{
			student{"Chelsea", 3},
			student{"Logan", 7},
		},
		expected: []string{"Chelsea", "Logan"},
	},
	{
		description: "Student not added to multiple grades in the roster",
		students: []student{
			student{"Blair", 2},
			student{"James", 2},
			student{"James", 3},
			student{"Paul", 3},
		},
		expected: []string{"Blair", "James", "Paul"},
	},
	{
		description: "Students are sorted by grades in the roster",
		students: []student{
			student{"Jim", 3},
			student{"Peter", 2},
			student{"Anna", 1},
		},
		expected: []string{"Anna", "Peter", "Jim"},
	},
	{
		description: "Students are sorted by name in the roster",
		students: []student{
			student{"Peter", 2},
			student{"Zoe", 2},
			student{"Alex", 2},
		},
		expected: []string{"Alex", "Peter", "Zoe"},
	},
	{
		description: "Students are sorted by grades and then by name in the roster",
		students: []student{
			student{"Peter", 2},
			student{"Anna", 1},
			student{"Barb", 1},
			student{"Zoe", 2},
			student{"Alex", 2},
			student{"Jim", 3},
			student{"Charlie", 1},
		},
		expected: []string{"Anna", "Barb", "Charlie", "Alex", "Peter", "Zoe", "Jim"},
	},
}

type testCaseGrade struct {
	description string
	students    []student
	grade       int
	expected    []string
}

var testCasesGrade = []testCaseGrade{
	{
		description: "Grade is empty if no students in the roster",
		students:    []student{},
		grade:       1,
		expected:    []string{},
	},
	{
		description: "Grade is empty if no students in that grade",
		students: []student{
			student{"Peter", 2},
			student{"Zoe", 2},
			student{"Alex", 2},
			student{"Jim", 3},
		},
		grade:    1,
		expected: []string{},
	},
	{
		description: "Student not added to same grade more than once",
		students: []student{
			student{"Blair", 2},
			student{"James", 2},
			student{"James", 2},
			student{"Paul", 2},
		},
		grade:    2,
		expected: []string{"Blair", "James", "Paul"},
	},
	{
		description: "Student not added to multiple grades",
		students: []student{
			student{"Blair", 2},
			student{"James", 2},
			student{"James", 3},
			student{"Paul", 3},
		},
		grade:    2,
		expected: []string{"Blair", "James"},
	},
	{
		description: "Student not added to other grade for multiple grades",
		students: []student{
			student{"Blair", 2},
			student{"James", 2},
			student{"James", 3},
			student{"Paul", 3},
		},
		grade:    3,
		expected: []string{"Paul"},
	},
	{
		description: "Students are sorted by name in a grade",
		students: []student{
			student{"Franklin", 5},
			student{"Bradley", 5},
			student{"Jeff", 1},
		},
		grade:    5,
		expected: []string{"Bradley", "Franklin"},
	},
}
