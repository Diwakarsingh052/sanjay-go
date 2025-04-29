package main

import "learn-go/slice"

func main() {

	x := []int{10, 20, 30, 40, 50, 60, 70}
	b := x[1:5]

	// cap of b calculation: - it starts from the 0 index of b till the end of the backing array,
	// in this case 6 is cap
	slice.Inspect("x", x)
	slice.Inspect("b", b) // [20 30 40 50]

	//below line would change x, b refers to the same backing array,
	//there is room to add two more elements to the existing backing array
	//so it would add the value to the backing array refer  by x
	b = append(b, 777)

	slice.Inspect("b", b)
	slice.Inspect("x", x)
	////b = append(b, 777, 888, 999) // this would create a new backing array for b,
	//	//it will not change the x slice

	b = append(b, 888, 999)

	slice.Inspect("b", b)
	slice.Inspect("x", x)
}
