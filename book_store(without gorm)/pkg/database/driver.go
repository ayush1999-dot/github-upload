package drivers

import (
	"database/sql"
	"log"
)

func Connectdb() *sql.DB {

	const DNS = "root:ayush@123@tcp(127.0.0.1:3306)/book-store"

	var err error
	var db *sql.DB
	db, err = sql.Open("mysql", DNS)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
