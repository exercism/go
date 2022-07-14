# About

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

## Buffered Channels

Channels default to being unbuffered, meaning they will only accept sends if a receiver is ready to receive.
But channels can be buffered.
Provide the buffer length as the second argument to make to initialize a buffered channel:

```go
ch := make(chan int, 150)
```

Channel will be blocked when the buffer is full and receives will block when the buffer is empty.

## Range and Close

A sender can `close` a channel to indicate that no more values will be sent.
Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after

```go
v, ok := <-ch
```

`ok` is `false` if there are no more values to receive and the channel is closed.

The loop `for i := range c` receives values from the channel repeatedly until it is closed.[^1]

### Notes[^1]:
~~~~exercism/note
Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
~~~~

~~~~exercism/note
Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.
~~~~

## Select

The `select` statement like switch-case in other languages, but for channels.
A `select` blocks until one of its cases can run, then it executes that case.
Order of run cases is random.

### Example
We want to brute force to find password of a lock, we use goroutines and channels to do this.

```go
package main

import "fmt"

func bruteForce(ch1, ch2, ch3 chan int) {
	lockPassword := 123456
	for {
		select {
		case x := <-ch1:
			if x == lockPassword {
				fmt.Println("Password found on channel 1:", x)
				return
			}
		case x := <-ch2:
			if x == lockPassword {
				fmt.Println("Password found on channel 2:", x)
				return
			}
		case x := <-ch3:
			if x == lockPassword {
				fmt.Println("Password found on channel 3:", x)
				return
			}
		}
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	go func() {
		for i := 0; i < 100000000; i++ {
			if i%3 == 1 {
				ch1 <- i
			} else if i%3 == 2 {
				ch2 <- i
			} else {
				ch3 <- i
			}
		}
	}()
	bruteForce(ch1, ch2, ch3)
}
```

[^1]: [Range and Close](https://go.dev/tour/concurrency/4)
