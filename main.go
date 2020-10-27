package main

import (
	"fmt"
	dbutils "github.com/Leonardo-Antonio/api-notes/DBUtils"
	"github.com/Leonardo-Antonio/api-notes/user"
)

func main() {
	pool := dbutils.GetConnection(dbutils.MYSQL)
	u := user.NewStorage(pool)
	err := u.CreateUser(user.User{
		LastName: "Nolasco",
		Email:    "AJNN@gmail.com",
		Password: "alexandra",
	})
	if err != nil {
		fmt.Println(err)
	}
}
