package lsproduct

import "fmt"

const testVersion = 3

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span < 0 {
		return 0, fmt.Errorf("span is negative: %d", span)
	}
	if len(digits) < span {
		return 0, fmt.Errorf("len(%s) < span: %d < %d", digits, len(digits), span)
	}
	v := make([]int64, len(digits))
	for i, r := range digits {
		if r < '0' || r > '9' {
			return 0, fmt.Errorf("input %q contains non-digits", digits)
		}
		v[i] = int64(r - '0')
	}
	maxsp := int64(0)
	for i, last := 0, len(v)-span+1; i < last; i++ {
		sp := int64(1)
		for _, d := range v[i : i+span] {
			sp *= d
		}
		if sp > maxsp {
			maxsp = sp
		}
	}
	return maxsp, nil
}
