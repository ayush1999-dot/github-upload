package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ayushshetty1999/pkg/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var err error

type Controller struct{}

func (c Controller) Getbooks(db *sql.DB) http.HandlerFunc {
	return func(wr http.ResponseWriter, r *http.Request) {
		wr.Header().Set("content-type", "application/json")
		row, err := db.Query("select *from books")

		if err != nil {
			log.Fatal(err)
		}
		for row.Next() {
			var book models.Book
			err = row.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
			if err != nil {
				log.Fatal(err)
			}
			json.NewEncoder(wr).Encode(&book)
		}

	}
}
func (c Controller) Getbook(db *sql.DB) http.HandlerFunc {
	return func(wr http.ResponseWriter, r *http.Request) {
		wr.Header().Set("content-type", "application/json")
		params := mux.Vars(r)
		check := params["id"]
		var book models.Book
		row1 := db.QueryRow("select *from books where id= ?", check)
		if err != nil {
			log.Fatal(err)
		}

		err = row1.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(wr).Encode(&book)
	}
}

func (c Controller) Addbook(db *sql.DB) http.HandlerFunc {
	return func(wr http.ResponseWriter, r *http.Request) {
		wr.Header().Set("content-type", "application/json")
		var book models.Book
		json.NewDecoder(r.Body).Decode(&book)
		stmt, err := db.Prepare("insert into books(title,author,year)values(?,?,?) ;")

		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		_, err = stmt.Exec(book.Title, book.Author, book.Year)

		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(wr).Encode(book)
	}
}
func (c Controller) Updatebook(db *sql.DB) http.HandlerFunc {
	return func(wr http.ResponseWriter, r *http.Request) {
		wr.Header().Set("content-type", "application/json")
		var book models.Book
		json.NewDecoder(r.Body).Decode(&book)
		param := mux.Vars(r)
		row, err := db.Prepare("update books set title=? ,author =? ,year=? where id =?")
		if err != nil {
			log.Fatal(err)
		}
		_, err = row.Exec(book.Title, book.Author, book.Year, param["id"])
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(wr).Encode(&book)
	}
}

func (c Controller) Deletebook(db *sql.DB) http.HandlerFunc {
	return func(wr http.ResponseWriter, r *http.Request) {
		wr.Header().Set("content-type", "application/json")

		params := mux.Vars(r)
		del, err := db.Prepare("delete from books where id =?")
		if err != nil {
			log.Fatal(err)
		}
		del.Exec(params["id"])

	}
}
