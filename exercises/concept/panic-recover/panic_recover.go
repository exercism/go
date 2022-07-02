package panicrecover

import "fmt"

// Add condition to panic
func AddPanic() {
	array := make([]int, 3)
	array[1] = 2
	array[2] = 4
	array[3] = 6
}

func panicHandler() {
	rec := recover()
	if rec != nil {
		fmt.Println("Recovering Panic", rec)
	}
}

// Recover from panic
func RecoverPanic(names []string, index int) {
	defer panicHandler()
	newName := names[index]
	fmt.Println(newName)
	//panic("Please add the recovering logic for the panic caused by Index Out of Bounds Error")
}

// Resolve error causing panic
func ResoveError() {
}
