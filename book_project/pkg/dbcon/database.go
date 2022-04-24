package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB
var err error

func Connect() *sqlx.DB {
	Db, err = sqlx.Connect("mysql", "root:ayush@123@(localhost:3306)/book_store")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("database connected successfully")
	}

	return Db
}
