package db

import "fmt"

type Conn struct {
	db string
}

func NewConn(db string) Conn {
	return Conn{db: db}
}

func (c Conn) Ping() {
	if c.db == "" {
		fmt.Println("db is empty, can't ping")
		return
	}
	fmt.Println("db is alive")
}
