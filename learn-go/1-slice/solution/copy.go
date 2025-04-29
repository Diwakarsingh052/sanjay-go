package main

import "fmt"

func main() {
	x := []int{10, 20, 30, 40, 50, 60, 70}

	// make allocates backing array for slice
	b := make([]int, len(x[1:5])) // new backing array // // make(type,len,cap)
	copy(b, x[1:5])

	b[0] = 999 // it will not affect the x slice anymore
	fmt.Println(b)
	fmt.Println(x)

}
