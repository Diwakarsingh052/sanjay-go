package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)

	get := make(chan string)
	post := make(chan string)
	put := make(chan string)

	wg.Add(3)
	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		get <- "get"
	}()

	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		post <- "post"
	}()

	go func() {
		defer wg.Done()
		put <- "put"
	}()

	//1st solution
	// run one goroutine for each channel
	//fmt.Println(<-get) // it is blocking call, value must be received first
	//fmt.Println(<-post)
	//fmt.Println(<-put)

	for i := 0; i < 3; i++ {
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
		}
	}

	wg.Wait()

}
