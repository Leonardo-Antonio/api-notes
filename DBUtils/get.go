package dbutils

import (
	"database/sql"
	"sync"
)

var (
	once sync.Once
	db *sql.DB
)

func GetConnection(typeDB string) *sql.DB {
	once.Do(func() {
		pool := newConnection(db)
		switch typeDB {
		case MSSQL:
			db = pool.Mssql()
		case MYSQL:
			db = pool.Mysql()
		case PSQL:
			db = pool.Psql()
		}
	})
	return db
}
