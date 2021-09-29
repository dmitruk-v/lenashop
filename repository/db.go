package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "user=postgres password=419155 dbname=lenashop sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}
