package main

import (
	"log"
	"net/http"

	"github.com/ayush1999-dot/go-project/pkg/handlers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func intitialization() {
	r := mux.NewRouter()
	r.HandleFunc("/Getuser", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/Getuser{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/Getuser", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/Getuser", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/Getuser", handlers.DeleteUser).Methods("Delete")
	log.Fatal(http.ListenAndServe(":8000", r))

}
func main() {
	//	http.HandleFunc("/productpage", handlers.Product)
	/*http.HandleFunc("/cart", handlers.Cart)
	http.HandleFunc("/Home", handlers.Homepage)
	http.HandleFunc("/productpage2", handlers.Productpage)
	http.ListenAndServe(":8080", nil)
	db, _ := sql.Open("mysql", "root:password@tcp(3306)/products")
	defer db.Close()*/
	handlers.Migrate()
	intitialization()

}
