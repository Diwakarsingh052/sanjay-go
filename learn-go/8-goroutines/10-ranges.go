package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wgWorker := new(sync.WaitGroup)
	ch := make(chan int, 10)
	// close the channel once the sender is finished
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			//we need to block our goroutine before closing the channel
			//because we want to make sure all the work
			// is done and finished
			// wgWorker waitgroup we are using to track number of worker goroutines
			wgWorker.Add(1)
			go func() {
				defer wgWorker.Done()
				ch <- i
			}()

		}
		// waiting until all the workers are not finished
		wgWorker.Wait()

		// closing when we are sure all the fanned out goroutines finished processing
		close(ch) // sends a signal to stop the range
		// close signals the range that no more values be sent and it can stop after receiving remaining values
		// once the channel is closed, we can't send more values to it
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch { // it would run infinitely, channel needs to be closed to stop this range
			fmt.Println("recv", v)
		}
	}()

	wg.Wait()

}
