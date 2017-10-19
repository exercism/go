package collatz

import "os"

var exit = os.Exit

func Steps(num int) (count int) {

	if num <= 0 {
		exit(1)
	}

	var n = num

	for n != 1 {
		count += 1
		switch n % 2 {
		case 0:
			n = n / 2
		case 1:
			n = n*3 + 1
		}
	}

	return
}
