package main

import "fmt"

func main() {
	optionalParam("abc", 10, 20, 30, 30)
	fmt.Println()
}

func optionalParam(s string, a ...int) {
	fmt.Println(s)
	fmt.Printf("%T \n", a)

}
