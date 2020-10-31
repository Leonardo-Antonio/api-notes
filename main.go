package main

import (
	dbutils "github.com/Leonardo-Antonio/api-notes/DBUtils"
	"github.com/Leonardo-Antonio/api-notes/notes"
	"github.com/Leonardo-Antonio/api-notes/type_notes"
	"github.com/Leonardo-Antonio/api-notes/user"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	pool := dbutils.GetConnection(dbutils.MYSQL)

	UserStorage := user.NewStorage(pool)
	TypeNotesStorage := type_notes.NewStorage(pool)
	NoteStorage := notes.NewStorage(pool)

	e := echo.New()
	e.Use(middleware.CORS())

	user.Router(UserStorage, e)
	type_notes.Router(TypeNotesStorage, e)
	notes.Router(NoteStorage, e)

	e.Logger.Fatal(e.Start(":8080"))
}
