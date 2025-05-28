package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func OpenFile(filename string) (*os.File, error) {
	f, err := os.Open(filename)
	if err != nil {
		//errors.Is can check if an error was wrapped inside the chain or not
		//  if an error was found in the chain, you now know what exactly went wrong
		// you might want to take some actions to fix the issue
		//or maybe just log the additional details
		if errors.Is(err, os.ErrNotExist) {
			//log.Println("attempting to create file")
			//log.Println("please create the file first then try to open")
			f, err := os.Create(filename)

			// if it still fails, we will return the original error
			if err != nil {
				return nil, err
			}

			// success case // errors.Is succeeded in identifying the root cause of the problems
			return f, nil

		}

		return nil, err
	}

	return f, nil
}

func main() {
	f, err := OpenFile("test.txt")
	if err != nil {
		log.Println(err)
		return
	}
	info, _ := f.Stat()
	fmt.Println("file size:", info.Name())
}
