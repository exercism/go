// +build detailed

package protein

import (
	"fmt"
	"testing"
)

func BenchmarkCodonDetailed(b *testing.B) {
	for _, test := range codonTestCases {
		b.Run(fmt.Sprintf("Codon%s", test.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FromCodon(test.input)
			}
		})
	}
}

func BenchmarkProteinDetailed(b *testing.B) {
	for _, test := range proteinTestCases {
		b.Run(fmt.Sprintf("Protein%s", test.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FromRNA(test.input)
			}
		})
	}
}
