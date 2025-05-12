package main

import (
	"fmt"
	"io"
)

// Polymorphism means that a piece of code changes its behavior depending on the
// concrete data it’s operating on // Tom Kurtz, Basic inventor

// "Don’t design with interfaces, discover them". - Rob Pike

// few exception- If you want to mock for testing,
// -or there are multiple types that need to implement interface immediately

// Bigger the interface weaker the abstraction // Rob Pike

//

type Reader interface {
	// interfaces are automatically implemented if method signature is same as of interfaces
	// only method signatures could be added to an interface

	Read(p []byte) (n int, err error)
	//write(b []byte) (int, error) // all methods must be implemented to implement the interface
}

//type Writer interface {
//	Read(p []byte) (n int, err error)
//	Write(p []byte) (n int, err error)
//}

type File struct {
	Name string
}

func (f File) Read(b []byte) (int, error) {
	fmt.Println("reading files and processing them", f.Name)
	return 0, nil
}

type IO struct {
	Name string
}

func (i IO) Read(b []byte) (int, error) {
	fmt.Println("reading and processing io ", i.Name)
	return 0, nil
}

func DoRead(r io.Reader, data []byte) {
	n, err := r.Read(data)
	_, _ = n, err

	fmt.Printf("type name %T\n", r)
	fmt.Println()
}

func main() {
	f := File{Name: "fs.txt"}
	i := IO{Name: "io1"}
	DoRead(f, make([]byte, 10))
	DoRead(i, make([]byte, 10))
}
