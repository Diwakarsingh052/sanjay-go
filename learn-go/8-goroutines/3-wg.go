package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go work(i, wg)
	}

	wg.Wait()
	fmt.Println("done")
}

func work(workId int, wg *sync.WaitGroup) {
	defer wg.Done()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("sub goroutine", "sub", workId)
	}()
	fmt.Println("work finished ", workId)

}
