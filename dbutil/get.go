package dbutil

import (
	"database/sql"
	"sync"
)

var (
	once sync.Once
	db   *sql.DB
)

func GetConnection(typeConn string) *sql.DB {
	once.Do(func() {
		con := newConnection(db)
		switch typeConn {
		case MYSQL:
			db = con.MYSQL()
		case MSSQL:
			db = con.MYSQL()
		}
	})
	return db
}
