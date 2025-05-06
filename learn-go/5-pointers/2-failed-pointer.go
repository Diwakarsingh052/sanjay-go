package main

import "fmt"

func main() {

	var p *int // nil // default value of a pointer is nil
	fmt.Println(&p)

	// after calling the update value p is still nil, as we never updated the pointer
	updateValue(p) // pass the value of p to updateValue func
	//fmt.Println(*p) // this line would panic,
}

func updateValue(p1 *int) {
	var x int = 10
	p1 = &x
	fmt.Println(&p1)
	fmt.Println(*p1)
}
