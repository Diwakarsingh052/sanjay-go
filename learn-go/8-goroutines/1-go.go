package main

import (
	"fmt"
	"time"
)

// concurrency is dealing with a lot of things at once
// parallelism is doing multiple things at once

func main() {
	//panic("some problem") // reveals goroutine id and name
	go hello() // this call would be concurrent, which means hello is a new goroutine

	fmt.Println("end of the main")
	time.Sleep(1 * time.Second)

}

func hello() {
	//time.Sleep(time.Second * 2)
	fmt.Println("Hello World")
}
