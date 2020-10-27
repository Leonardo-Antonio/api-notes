package dbutils

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type connection struct {
	db *sql.DB
}

func newConnection(db *sql.DB) *connection {
	return &connection{db}
}

func (c *connection) Mysql() *sql.DB {
	db, err := sql.Open(
		"mysql",
		"leo:chester@tcp(localhost:3306)/DB_Notes",
		)
	if err != nil {
		log.Fatalf("Error: %v\nFailed to connect to the database", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Error: %v\nCould not ping", err)
	}
	c.db = db
	return c.db
}

func (c *connection) Mssql() *sql.DB {
	db, err := sql.Open(
		"mysql",
		"leo:chester@tcp(localhost:3306)/db_notes")
	if err != nil {
		log.Fatalf("Error: %v\nFailed to connect to the database", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Error: %v\nCould not ping", err)
	}
	c.db = db
	return c.db
}

func (c *connection) Psql() *sql.DB {
	db, err := sql.Open(
		"mysql",
		"leo:chester@tcp(localhost:3306)/db_notes")
	if err != nil {
		log.Fatalf("Error: %v\nFailed to connect to the database", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Error: %v\nCould not ping", err)
	}
	c.db = db
	return c.db
}