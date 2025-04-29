package db

// you can't have two packages one folder

import "fmt"

var db = ""

func CreateConn(name string) {
	db = name
}
func GetConn() string {
	return db
}
func fetch() {
	fmt.Println("fetching the records from the db", db)
}
