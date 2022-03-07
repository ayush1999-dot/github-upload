package main

import (
	"database/sql"
	"net/http"

	"github.com/ayush1999-dot/go-project/pkg/handlers"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	//	http.HandleFunc("/productpage", handlers.Product)
	http.HandleFunc("/productpage/buy", handlers.Cart)
	http.ListenAndServe(":8080", nil)
	db, _ := sql.Open("mysql", "root:password@tcp(3306)/products")
	defer db.Close()

}
