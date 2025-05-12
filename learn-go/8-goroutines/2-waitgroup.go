package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//var wg = &sync.WaitGroup{} //

	// using waitgroup keep track of goroutines
	wg := new(sync.WaitGroup)

	// waitgroup counter represents number of goroutine we are running
	wg.Add(1) //
	go func(i int) {
		// giving a guarantee that even
		// in case of panic this would decrement the counter
		defer wg.Done() // decrement the counter by 1
		time.Sleep(time.Second)
		fmt.Println("hello world")

	}(10)

	fmt.Println("some work in main going on")
	wg.Wait() // wait until the counter is not 0
	fmt.Println("end of the main")
}
