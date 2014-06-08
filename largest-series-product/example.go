package lsproduct

import "fmt"

func LargestSeriesProduct(digits string, span int) (int, error) {
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
	maxsp := 1
	for i, last := 0, len(v)-span+1; i < last; i++ {
		sp := 1
		for _, d := range v[i : i+span] {
			sp *= d
		}
		if sp > maxsp {
			maxsp = sp
		}
	}
	return maxsp, nil
}
