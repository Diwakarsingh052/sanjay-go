package main

import "fmt"

// https://go.dev/ref/spec#Appending_and_copying_slices
func main() {
	a := []int{10, 20, 30, 40, 50}
	// %p prints the address of the first element
	fmt.Printf("before append %p\n", a)
	fmt.Println(len(a), cap(a))

	//a[4] = 1000 // update operation

	a = append(a, 88, 99)
	fmt.Printf("after append %p\n", a)
	fmt.Println(len(a), cap(a))

	a = append(a, 110)

	fmt.Printf("after second append %p\n", a)
	fmt.Println(len(a), cap(a))

	fmt.Println(a)
}
