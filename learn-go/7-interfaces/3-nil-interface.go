package main

import (
	"database/sql"
	"fmt"
	"log"
)

type DB interface {
	ReadAll()
}

type Conn struct {
	db   *sql.DB // default value would be nil
	data string
}

func New(db *sql.DB) *Conn {
	if db == nil {
		log.Println("db is nil, please provide a valid connection")
		return nil
	}
	return &Conn{db: db}
}

func (c *Conn) ReadAll() {

	if c == nil {
		return
	}
	if c.db == nil {
		return
	}

	c.data = "abc"
	c.db.Ping()
	fmt.Println("reading some data", c.data)

}

func main() {

	//var c *conn // nil
	c := new(Conn)
	var db DB = c

	if db == nil {
		log.Println("db is nil")
		return
	}
	fmt.Println("calling the read all method")
	db.ReadAll()
}
