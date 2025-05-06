package main

import (
	"fmt"
)

// creating a new type operation,
type operation func(int, int) int

type money int

// abc is an alias not a new type
type abc = int

// operate func can accept function in op parameter,
// the function signature we are passing should match to op parameter type
func operateV2(op operation, x, y int) {
	fmt.Println("inside op ")
	fmt.Println(op(x, y))
}

// datatype of func -> func(args)returnType
func addV2() operation {
	fmt.Println("inside add outer")
	return func(a int, b int) int {
		fmt.Println("inside add inner returned func")
		return a + b
	}
}

func subV2(a, b int) int {
	return a - b
}

func main() {
	operate(add(), 10, 20)
	//operate(sub, 90, 10)

}
