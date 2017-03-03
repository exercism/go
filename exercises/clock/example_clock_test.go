package clock

import "fmt"

func ExampleClock() {
	// Create a clock
	clock1 := New(10, 30)
	fmt.Println(clock1.String())

	// Add 30 minutes to it
	clock1 = clock1.Add(30)
	fmt.Println(clock1.String())

	// Subtract an hour and a half from it
	clock1 = clock1.Add(-90)
	fmt.Println(clock1.String())

	// Create a second clock
	clock2 := New(9, 30)
	fmt.Println(clock2.String())

	// Are the clocks equal?
	fmt.Println(clock2 == clock1)

	// Change the second clock
	clock2 = clock2.Add(30)
	fmt.Println(clock2.String())

	// Are the clocks equal now?
	fmt.Println(clock2 == clock1)

	// Output:
	// 10:30
	// 11:00
	// 09:30
	// 09:30
	// true
	// 10:00
	// false
}
