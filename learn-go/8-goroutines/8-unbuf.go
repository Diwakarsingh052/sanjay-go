package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 4; i++ {

			ch <- i

			fmt.Println("sent", i)
		}
		// blocked until there is a recv

	}()

	for i := 1; i <= 4; i++ {
		fmt.Println(<-ch)

	}

	wg.Wait()
}
