package main

import (
	"fmt"
	"sync"
)

var x int = 1

// maps are not safe for concurrent access, use mutex for safe operations
var user = map[int]string{
	1: "John",
}

func main() {

	wg := new(sync.WaitGroup)
	m := new(sync.Mutex)
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go UpdateX(i, wg, m)
		go UpdateVariable(i, wg)
	}

	wg.Wait()

}

func UpdateX(val int, wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()
	// critical section
	// this is the place where we access the shared resource

	// when a goroutine acquires a lock, another goroutine can't access the critical section
	// until the lock is not released

	// data race situations
	//	- at least one concurrent write and n number of reads
	//	- n number of concurrent writes
	// 	- n number of concurrent writes and n number of concurrent reads
	// 	Note - Data race doesn't happen if there are only concurrent reads

	m.Lock()
	defer m.Unlock()
	//the mutex is not locking a specific variable its locking a block of code
	x = val
	//time.Sleep(1 * time.Second)
	fmt.Println(x)
	// locking till the end of the function
}

func UpdateVariable(val int, wg *sync.WaitGroup) {
	defer wg.Done()
	var abc int
	abc = val
	fmt.Println(abc)
}
