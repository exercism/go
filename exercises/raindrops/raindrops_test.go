package raindrops

import "testing"

const targetTestVersion = 2

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
}

func TestConvert(t *testing.T) {
	for _, test := range tests {
		if actual := Convert(test.input); actual != test.expected {
			t.Errorf("Convert(%d) = %q, expected %q.",
				test.input, actual, test.expected)
		}
	}
}

func BenchmarkConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			Convert(test.input)
		}
	}
}
