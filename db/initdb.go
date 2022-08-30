package db

import (
	"log"

	"github.com/jmoiron/sqlx"
)

var dsn = "postgres://vijay:12345@localhost:5432/ginrest?sslmode=disable"

func InitDb() *sqlx.DB {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("\n\nconnection successful\n\n")
	}
	return db
}
