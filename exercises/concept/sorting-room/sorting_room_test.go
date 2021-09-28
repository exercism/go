package sorting

import "testing"

func TestDescribeNumber(t *testing.T) {
	tests := []struct {
		description string
		input       float64
		want        string
	}{
		{
			desc:  "Describe 4.1",
			input: 4.1,
			want:  "This is the number 4.1",
		},
		{
			desc:  "Describe -3.2",
			input: -3.2,
			want:  "This is the number -3.2",
		},
		{
			desc:  "Pads to single decimal place",
			input: 4.0,
			want:  "This is the number 4.0",
		},
		{
			desc:  "Truncates to single decimal place",
			input: 7.11,
			want:  "This is the number 7.1",
		},
	}
	for _, test := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if got := DescribeNumber(test.input); got != test.want {
				t.Errorf("DescribeNumber(%v) = %v; want %v", test.input, got, test.want)
			}
		})
	}
}
