# Introduction

[Goroutines](https://go.dev/tour/concurrency/1) are a way to execute concurrent code in Go.
Goroutines are works independently, if they need to communicate, Go preferred to use channels.
Create a new channel with `make(chan val-type)`.
Channels are typed by the values they convey.

```go
ch := make(chan int) // create new channel
ch <- value          // Send value(same as channel type) to channel ch.
value := <-ch        // Receive from ch and assign value to value variable.
```

By default, sends and receives block wait until the other side is ready.
This allows goroutines to synchronize without explicit locks or condition variables.

Example of channel usage, if we want to findout the value exists in the slice, we can split the slice into two parts, find separately in other goroutines, and then send the result to the main goroutine.

```go
package main

import "fmt"

func exist(slice []int, x int, resultChan chan bool) {
	found := false
	for _, v := range slice {
		if v == x {
			found = true
            break
		}
	}
	resultChan <- found // send found to channel
}

func main() {
	slice := []int{1, 7, 8, -9, 12, -1, 13, -7}
	lookForValue := 8

	resultChan := make(chan bool)
	go exist(slice[:len(slice)/2], lookForValue, resultChan)
	go exist(slice[len(slice)/2:], lookForValue, resultChan)
	firstHalf, secondHalf := <-resultChan, <-resultChan // receive 2 values from channel

	fmt.Println(firstHalf || secondHalf)
}
```
