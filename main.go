package main

import (
	dbutils "github.com/Leonardo-Antonio/api-notes/DBUtils"
	"github.com/Leonardo-Antonio/api-notes/type_notes"
	"github.com/Leonardo-Antonio/api-notes/user"
	"github.com/labstack/echo"
)

func main() {
	pool := dbutils.GetConnection(dbutils.MYSQL)
	UserStorage := user.NewStorage(pool)
	TypeNotesStorage := type_notes.NewStorage(pool)

	e := echo.New()

	user.Router(UserStorage, e)
	type_notes.Router(TypeNotesStorage, e)
	e.Logger.Fatal(e.Start(":8080"))
}