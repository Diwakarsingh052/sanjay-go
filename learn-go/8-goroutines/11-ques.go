package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	m := new(sync.Mutex)
	//wgWorker := new(sync.WaitGroup)
	counter := 0
	numRoutines := 1000000

	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			m.Lock()
			counter++
			m.Unlock()
		}()

	}
	wg.Wait()
	fmt.Println("Final Counter:", counter) // ??
}
