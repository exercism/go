package railfencecipher

import "testing"

func TestEncode(t *testing.T) {
	for _, tc := range encodeTests {
		t.Run(tc.description, func(t *testing.T) {
			if actual := Encode(tc.message, tc.rails); actual != tc.expected {
				t.Fatalf("Encode(%q, %d)\ngot: %q\nwant: %q", tc.message, tc.rails, actual, tc.expected)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	for _, tc := range decodeTests {
		t.Run(tc.description, func(t *testing.T) {
			if actual := Decode(tc.message, tc.rails); actual != tc.expected {
				t.Fatalf("Decode(%q, %d)\ngot: %q\nwant: %q", tc.message, tc.rails, actual, tc.expected)
			}
		})
	}
}

func BenchmarkEncode(b *testing.B) {
	for range b.N {
		for _, test := range encodeTests {
			Encode(test.message, test.rails)
		}
	}
}

func BenchmarkDecode(b *testing.B) {
	for range b.N {
		for _, test := range decodeTests {
			Decode(test.message, test.rails)
		}
	}
}
