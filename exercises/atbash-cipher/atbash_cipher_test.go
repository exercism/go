package atbash

import "testing"

func TestAtbash(t *testing.T) {
	for _, test := range tests {
		actual := Atbash(test.s)
		if actual != test.expected {
			t.Errorf("Atbash(%s): expected %s, actual %s", test.s, test.expected, actual)
		}
	}
}

func BenchmarkAtbash(b *testing.B) {
	b.StopTimer()
	for _, test := range tests {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			Atbash(test.s)
		}

		b.StopTimer()
	}
}
