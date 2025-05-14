package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	wgWorker := new(sync.WaitGroup)

	get := make(chan string)
	post := make(chan string)
	put := make(chan string)
	done := make(chan struct{})

	wgWorker.Add(3)
	go func() {
		defer wgWorker.Done()
		get <- "get"
	}()

	go func() {
		defer wgWorker.Done()
		post <- "post"
	}()

	go func() {
		defer wgWorker.Done()
		put <- "put"
		put <- "p1"
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		wgWorker.Wait()
		close(done) // closing is a send signal, which can be received by a recv
		// some kind of signal to select
	}()

	//1st solution
	// run one goroutine for each channel
	// run a range inside the goroutine
	// 3 goroutines, 3 ranges

	//fmt.Println(<-get) // it is blocking call, value must be received first
	//fmt.Println(<-post)
	//fmt.Println(<-put)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {

			//whichever case is not blocking exec that first
			//whichever case is ready first, exec that.
			// possible cases are chan recv , send , default
			case msg := <-get:
				fmt.Println(msg)
			case msg := <-post:
				fmt.Println(msg)
			case msg := <-put:
				fmt.Println(msg)
				time.Sleep(2 * time.Second)
			case <-done: // once a channel is close, we would get a signal here
				fmt.Println("all values received, closing the select")
				return
			}
		}
	}()
	wg.Wait()

}
