package lsproduct

import "fmt"

const TestVersion = 1

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if len(digits) < span {
		return 0, fmt.Errorf("len(%s) < span: %d < %d", digits, len(digits), span)
	}
	v := make([]int, len(digits))
	for i, r := range digits {
		if r < '0' || r > '9' {
			return 0, fmt.Errorf("input %q contains non-digits", digits)
		}
		v[i] = int(r - '0')
	}
	maxsp := int64(1)
	for i, last := 0, len(v)-span+1; i < last; i++ {
		sp := int64(1)
		for _, d := range v[i : i+span] {
			sp *= int64(d)
		}
		if sp > maxsp {
			maxsp = sp
		}
	}
	return maxsp, nil
}
