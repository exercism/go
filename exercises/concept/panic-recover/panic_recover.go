package panicrecover

import "fmt"

// Add condition to panic
func AddPanic(name []string, index int) {
	val := name[index]
	fmt.Println(val)
}

func RecoverPanic() {
	rec := recover()
	if rec != nil {
		fmt.Println("Error due to which function ran into panic: ", rec)
	}
}

// Resolve error causing panic
func ResoveError() {
}
