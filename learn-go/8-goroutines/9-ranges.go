package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	ch := make(chan int, 10)
	// close the channel once the sender is finished
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch) // sends a signal to stop the range
		// close signals the range that no more values be sent and it can stop after receiving remaining values
		// once the channel is closed, we can't send more values to it
	}()

	time.Sleep(time.Second)
	for v := range ch { // it would run infinitely, channel needs to be closed to stop this range
		fmt.Println("recv", v)
	}

	wg.Wait()
}
