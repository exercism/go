package pangram

import (
	"testing"
)

const targetTestVersion = 1

type testCase struct {
	input         string
	expected      bool
	failureReason string
}

var testCases = []testCase{
	{"", false, "sentence empty"},
	{"The quick brown fox jumps over the lazy dog", true, ""},
	{"a quick movement of the enemy will jeopardize five gunboats", false, "missing character 'x'"},
	{"the quick brown fish jumps over the lazy dog", false, "another missing character 'x'"},
	{"the 1 quick brown fox jumps over the 2 lazy dogs", true, ""},
	{"7h3 qu1ck brown fox jumps ov3r 7h3 lazy dog", false, "missing letters replaced by numbers"},
	{"\"Five quacking Zephyrs jolt my wax bed.\"", true, ""},
	{"Victor jagt zwölf Boxkämpfer quer über den großen Sylter Deich.", true, ""},
	{"Широкая электрификация южных губерний даст мощный толчок подъёму сельского хозяйства.", false, "Panagram in alphabet other than ASCII"},
}

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v.", testVersion, targetTestVersion)
	}
}

func TestPangram(t *testing.T) {
	for _, test := range testCases {
		actual := IsPangram(test.input)
		if actual != test.expected {
			t.Errorf("Pangram test [%s], expected [%t], actual [%t]", test.input, test.expected, actual)
			if !test.expected {
				t.Logf("[%s] should not be a pangram because : %s\n", test.input, test.failureReason)
			}
		}
	}
}

func BenchmarkPangram(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			IsPangram(test.input)
		}
	}
}
