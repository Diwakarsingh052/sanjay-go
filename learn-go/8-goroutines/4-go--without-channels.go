package main

import (
	"fmt"
	"sync"
	"time"
)

func helloV2() string {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello V2")
	return "Hello"

}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		helloV2()
	}()

	wg.Wait()

}
