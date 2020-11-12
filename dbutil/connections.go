package dbutil

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

func (c *connection) MYSQL() *sql.DB {
	var err error
	c.db, err = sql.Open(
		"mysql",
		"leo:chester@tcp(localhost:3306)/db_profile")
	if err != nil {
		log.Fatalf("error trying to connect to database\n\v", err)
	}
	if err = c.db.Ping(); err != nil {
		log.Fatalf("error in Ping\n\v", err)
	}
	return c.db
}
