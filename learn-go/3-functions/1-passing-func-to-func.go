package main

import (
	"fmt"
)

// operate func can accept function in op parameter,
// the function signature we are passing should match to op parameter type
func operate(op func(int, int) int, x, y int) {
	fmt.Println("inside op ")
	fmt.Println(op(x, y))
}

// datatype of func -> func(args)returnType
func add() func(int, int) int {
	fmt.Println("inside add outer")
	return func(a int, b int) int {
		fmt.Println("inside add inner returned func")
		return a + b
	}
}

func sub(a, b int) int {
	return a - b
}

func main() {
	operate(add(), 10, 20)
	//operate(sub, 90, 10)

}
