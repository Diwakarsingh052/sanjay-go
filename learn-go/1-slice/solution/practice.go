package main

import "fmt"

func main() {

	x := []int{10, 20, 30, 40, 50} // l = 5, c=5
	x = updateSlice(x, 0, 999)
	fmt.Println(x)
}

func updateSlice(s []int, index, val int) []int {
	// if you want to update the value ,
	//then passing the reference of existing backing array is fine
	s[index] = val

	s = append(s, 110)
	fmt.Println("update slice", s)
	return s // make sure to return the updated slice if using append
}

// if a function is not suppose to modify the original slice
// then no update or append operation should be performed directly
// make sure you are working on a copy if need to update or append
