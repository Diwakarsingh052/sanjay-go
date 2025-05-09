package main

import (
	"fmt"
	"learn-go/db"
)

func main() {

	conn := db.NewConn("mysql")
	conn.Ping()
	fmt.Println(conn)

	conn1 := db.NewConn("pg")
	conn1.Ping()

	//log.New()
	//os.OpenFile()
}
