package models

import "fmt"

type Person struct {
	Name    string // fields of struct
	Age     int
	Address string
}

var Login struct {
	Name string
}

func (p Person) Print() {
	fmt.Println("Name:", p.Name)
}
