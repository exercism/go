package atbashcipher

import "testing"

func TestAtbash(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Atbash(tc.phrase)
			if actual != tc.expected {
				t.Errorf("Atbash(%q): expected %q, actual %q", tc.phrase, tc.expected, actual)
			}
		})
	}
}

func BenchmarkAtbash(b *testing.B) {
	for range b.N {
		for _, test := range testCases {
			Atbash(test.phrase)
		}
	}
}
