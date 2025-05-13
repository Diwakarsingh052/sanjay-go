package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for {
			time.Sleep(5 * time.Second)
		}
	}()

	fmt.Println("created a channel")
	<-ch

}
