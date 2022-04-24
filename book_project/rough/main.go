package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Book struct {
	ID         int    `json:"id"`
	Book_name  string `json:"book_name"`
	Author_Id  int    `json:"author_id"`
	Book_url   string `json:"book_url"`
	Is_deleted string `json:"is_deleted"`
}

var Db *sqlx.DB
var err error

func main() {
	Db, err = sqlx.Connect("mysql", "root:ayush@123@(localhost:3306)/book_store")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("database connected successfully")
	}
	defer Db.Close()
	/* book := Book{
		Book_name: "atomic habits",
		Author_Id: 101,
		Book_url:  "wwwewas.com",
	} */
     
	book := []Book{}
	Db.Select(&book, "SELECT * FROM books  ")
	fmt.Println(book)
	//Db.NamedExec("INSERT INTO books(Book_name, Author_Id, Book_url) VALUES (:book_name, :author_id, :book_url)", book)

}

/*
{
"book_name":"classmate",
"author_id":108,
"author_name":"life"
"book_url":"www.classmate.com",
"quantity_in_stock":20,
"price":200
}
  {

            "book_name": "siddhartha",
            "book_url": "www.siddtharta.com"
        }
*/
