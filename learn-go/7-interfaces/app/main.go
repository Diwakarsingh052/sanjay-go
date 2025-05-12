package main

import (
	"app/internal/stores"
	"app/internal/stores/models"
	"app/internal/stores/mysql"
	"app/internal/stores/postgres"
)

func main() {

	u1 := models.User{Id: 1, Name: "sandra"}
	u2 := models.User{Id: 2, Name: "komal"}
	m := mysql.NewConn()
	p := postgress.NewConn()

	AddUser(m, u1)
	AddUser(p, u2)

}

func AddUser(d stores.DataBase, u models.User) {
	d.Create(u)
}
