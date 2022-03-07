package main

import (
	"database/sql"
	"net/http"

	drivers "github.com/ayushshetty1999/pkg/database"
	handlers "github.com/ayushshetty1999/pkg/handler"
	"github.com/gorilla/mux"
)

func main() {
	var db *sql.DB = drivers.Connectdb()
	controller := handlers.Controller{}

	r := mux.NewRouter()
	r.HandleFunc("/books", controller.Getbooks(db)).Methods("GET")
	r.HandleFunc("/books/{id}", controller.Getbook(db)).Methods("GET")
	r.HandleFunc("/books", controller.Addbook(db)).Methods("POST")
	r.HandleFunc("/books/{id}", controller.Updatebook(db)).Methods("PUT")
	r.HandleFunc("/books/{id}", controller.Deletebook(db)).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}
