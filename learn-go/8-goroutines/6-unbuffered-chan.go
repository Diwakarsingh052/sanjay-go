package main

import (
	"fmt"
	"sync"
	"time"
)

// https://go.dev/ref/spec#Send_statements

// A send on an unbuffered channel can proceed if a receiver is ready.
// send will block until there is no recv
// channels are only meant to be used in concurrent programming

//A send on a buffered channel can proceed if there is room in the buffer.

// the receiver is always blocking whether it is buffered or unbuffered

func main() {
	wg := new(sync.WaitGroup)
	ch := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- 10
		fmt.Println("value sent to the channel")
	}()

	time.Sleep(3 * time.Second)
	fmt.Println(<-ch) //recv // this would block if no sender is present,
	//and another goroutine from the queue would be picked up
	//which is sender goroutine in this case
	//communication completes, and we get 10 on the screen
	wg.Wait()

}
