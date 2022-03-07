package main

import (
	"database/sql"
	"net/http"

	"github.com/ayush1999-dot/go-project/pkg/handlers"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//	http.HandleFunc("/productpage", handlers.Product)
	http.HandleFunc("/cart", handlers.Cart)
	http.HandleFunc("/Home", handlers.Homepage)
	http.HandleFunc("/productpage2", handlers.Productpage)
	http.ListenAndServe(":8080", nil)
	db, _ := sql.Open("mysql", "root:password@tcp(3306)/products")
	defer db.Close()

}
