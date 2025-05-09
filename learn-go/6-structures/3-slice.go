package main

import (
	"fmt"
)

type Author struct {
	Name  string
	Books []string
}

func (a *Author) AddBook(book string) {
	a.Books = append(a.Books, book)

}

// AppendBooksToAuthor accepts the Author struct and appends values to the Books slice
func AppendBooksToAuthor(author *Author, booksToAdd []string) {
	author.Books = append(author.Books, booksToAdd...)
}

// AppendToBooks accepts the Books slice and appends some more books
func AppendToBooks(books *[]string, booksToAdd []string) {

	// ... unpacking a slice
	*books = append(*books, booksToAdd...) // Dereferencing the pointer
}

func (a *Author) PrintDetails() {
	fmt.Printf("Author FirstName: %s\nBooks: %v\n", (*a).Name, a.Books)
}

func main() {
	a := &Author{} // below line have same affect
	a1 := new(Author)
	_, _ = a, a1

	author := Author{Name: "John Doe", Books: []string{"Book 1", "Book 2"}}
	//var a *Author //
	author.AddBook("book 3")
	author.PrintDetails()

	// Append more books using a function that accepts the struct
	AppendBooksToAuthor(&author, []string{"Book 4", "Book 5"})
	author.PrintDetails()

	// Append more books using a function that accepts the Books slice
	AppendToBooks(&author.Books, []string{"Book 6", "Book 7"})
	author.PrintDetails()

}
