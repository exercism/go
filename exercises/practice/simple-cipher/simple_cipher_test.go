package cipher

import (
	"fmt"
	"testing"
)

// type for testing cipher encoding alone, without requiring any text prep.
type prepped struct {
	pt string // prepped text == decoded plain text
	ct string // cipher text
}

// +3? -3?  Positions vary.  Your code just has to pass the tests.
var caesarPrepped = []prepped{
	{"iamapandabear", "ldpdsdqgdehdu"},
	{"programmingisawesome", "surjudpplqjlvdzhvrph"},
	{"todayisholiday", "wrgdblvkrolgdb"},
	{"venividivici", "yhqlylglylfl"},
}

func TestCaesarPrepped(t *testing.T) {
	c := NewCaesar()
	for _, test := range caesarPrepped {
		if enc := c.Encode(test.pt); enc != test.ct {
			t.Fatalf("Caesar Encode(%q) = %q, want %q.", test.pt, enc, test.ct)
		}
	}
}

// type for testing implementations of the Cipher interface
type cipherTest struct {
	source string // source text, any UTF-8
	cipher string // cipher text, result of Encode(st)
	plain  string // decoded plain text, result of Decode(ct)
}

var caesarTests = []cipherTest{
	{"Go, go, gophers", "jrjrjrskhuv", "gogogophers"},
	{"I am a panda bear.", "ldpdsdqgdehdu", "iamapandabear"},
	{"Programming is AWESOME!", "surjudpplqjlvdzhvrph", "programmingisawesome"},
	{"today is holiday", "wrgdblvkrolgdb", "todayisholiday"},
	{"Twas the night before Christmas",
		"wzdvwkhqljkwehiruhfkulvwpdv",
		"twasthenightbeforechristmas"},
	{"venividivici", "yhqlylglylfl", "venividivici"},
	{" -- @#!", "", ""},
	{"", "", ""},
}

func TestCaesar(t *testing.T) {
	testCipher("Caesar", NewCaesar(), caesarTests, t)
}

func testCipher(name string, c Cipher, tests []cipherTest, t *testing.T) {
	for _, test := range tests {
		if enc := c.Encode(test.source); enc != test.cipher {
			t.Fatalf("%s Encode(%q) = %q, want %q.",
				name, test.plain, enc, test.cipher)
		}
		if dec := c.Decode(test.cipher); dec != test.plain {
			t.Fatalf("%s Decode(%q) = %q, want %q.",
				name, test.cipher, dec, test.plain)
		}
	}
}

var NSATests = []cipherTest{
	{"THE ENEMY IS NEAR", "qebbkbjvfpkbxo", "theenemyisnear"},
	{"SPIES SEND SECRET MESSAGES",
		"pmfbppbkapbzobqjbppxdbp",
		"spiessendsecretmessages"},
	{"THOMAS JEFFERSON DESIGNED A SUBSTITUTION CIPHER",
		"qeljxpgbccboplkabpfdkbaxprypqfqrqflkzfmebo",
		"thomasjeffersondesignedasubstitutioncipher"},
	{"the quick brown fox jumps over the lazy dog",
		"qebnrfzhyoltkclugrjmplsboqebixwvald",
		"thequickbrownfoxjumpsoverthelazydog"},
}

func TestShift(t *testing.T) {
	// test shift(3) against Caesar cases.
	c := NewShift(3)
	if c == nil {
		t.Fatal("NewShift(3) returned nil, want non-nil Cipher")
	}
	testCipher("Shift(3)", c, caesarTests, t)

	// NSA and WP say Caesar uses shift of -3
	c = NewShift(-3)
	if c == nil {
		t.Fatal("NewShift(-3) returned nil, want non-nil Cipher")
	}
	testCipher("Shift(-3)", c, NSATests, t)

	// invalid shifts
	for _, s := range []int{-27, -26, 0, 26, 27} {
		if NewShift(s) != nil {
			t.Fatalf("NewShift(%d) returned non-nil, "+
				"Want nil return for invalid argument.", s)
		}
	}
}

var vtests = []struct {
	key   string
	tests []cipherTest
}{
	{"lemon", []cipherTest{{"ATTACKATDAWN", "lxfopvefrnhr", "attackatdawn"}}},
	{"abcdefghij", []cipherTest{
		{"aaaaaaaaaa", "abcdefghij", "aaaaaaaaaa"},
		{"zzzzzzzzzz", "zabcdefghi", "zzzzzzzzzz"},
	}},
	{"iamapandabear", []cipherTest{
		{"I am a panda bear.", "qayaeaagaciai", "iamapandabear"},
	}},
	{"duxrceqyaimciuucnelkeoxjhdyduu", []cipherTest{
		{"Diffie Hellman", "gccwkixcltycv", "diffiehellman"},
	}},
	{"qgbvno", []cipherTest{
		{"cof-FEE, 123!", "sugars", "coffee"},
	}},
}

func TestVigenere(t *testing.T) {
	for _, test := range vtests {
		v := NewVigenere(test.key)
		if v == nil {
			t.Fatalf("NewVigenere(%q) returned nil, want non-nil Cipher",
				test.key)
		}
		testCipher(fmt.Sprintf("Vigenere(%q)", test.key), v, test.tests, t)
	}

	// invalid keys
	for _, k := range []string{"", "a", "aa", "no way", "CAT", "3", "and,"} {
		if NewVigenere(k) != nil {
			t.Fatalf("NewVigenere(%q) returned non-nil, "+
				"Want nil return for invalid argument.", k)
		}
	}
}

// Benchmark combined time to run all tests.
// Note other ciphers test different data; times cannot be compared.
func BenchmarkEncodeCaesar(b *testing.B) {
	c := NewCaesar()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, test := range caesarTests {
			c.Encode(test.source)
		}
	}
}

func BenchmarkDecodeCaesar(b *testing.B) {
	c := NewCaesar()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, test := range caesarTests {
			c.Decode(test.cipher)
		}
	}
}

func BenchmarkNewShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for s := -27; s <= 27; s++ {
			NewShift(s)
		}
	}
}

func BenchmarkEncodeShift(b *testing.B) {
	s := NewShift(5)
	all := append(caesarTests, NSATests...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, test := range all {
			s.Encode(test.source)
		}
	}
}

func BenchmarkDecodeShift(b *testing.B) {
	s := NewShift(5)
	all := append(caesarTests, NSATests...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, test := range all {
			s.Decode(test.cipher)
		}
	}
}

func BenchmarkNewVigenere(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range vtests {
			NewVigenere(test.key)
		}
	}
}

func BenchmarkEncVigenere(b *testing.B) {
	v := make([]Cipher, len(vtests))
	for i, test := range vtests {
		v[i] = NewVigenere(test.key)
		if v[i] == nil {
			b.Skip("Benchmark requires valid Vigenere test cases")
		}
	}
	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		for i, test := range vtests {
			vi := v[i]
			for _, test := range test.tests {
				vi.Encode(test.source)
			}
		}
	}
}

func BenchmarkDecVigenere(b *testing.B) {
	v := make([]Cipher, len(vtests))
	for i, test := range vtests {
		v[i] = NewVigenere(test.key)
		if v[i] == nil {
			b.Skip("Benchmark requires valid Vigenere test cases")
		}
	}
	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		for i, test := range vtests {
			vi := v[i]
			for _, test := range test.tests {
				vi.Decode(test.cipher)
			}
		}
	}
}
