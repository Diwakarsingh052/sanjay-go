package main

import "learn-go/db"

// package name must not start with a number
// package naming must be simple
// avoid package names like util, helpers, common
// package naming trick: think about what functionality it provides
// if multiple words in package, then don't use - , or _, use all smallcase letters

// Exporting
// We need to make the first letter as uppercase to export the function
//E.g. Inspect()

// folder name and the package name must be the same

func main() {
	//db.DB = "postgres"
	db.CreateConn("mysql")
	db.CreateConn("postgres")
	db.GetConn()
	db.PrintRecords()

}
