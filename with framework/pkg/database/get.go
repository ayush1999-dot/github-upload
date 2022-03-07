package utilities

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type detail struct {
	Id           int
	Productname  string
	Image        string
	Priceold     int
	Price        int
	Connectivity int
}

var map_1 map[string]string = map[string]string{
	"00": "select * from m_wireless",
	"01": "select * from m_wired",
	"10": "select * from k_wireless",
	"11": "select * from k_wired",
}

func Getdata(home string) []detail {
	var results *sql.Rows
	db, err := sql.Open("mysql", "root:ayush@123@tcp(127.0.0.1:3306)/product")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	for i, j := range map_1 {
		if home == i {
			results, err = db.Query(j)
		}
	}

	if err != nil {
		log.Fatal(err)
	}
	var p1 detail
	var p2 []detail
	for results.Next() {

		err = results.Scan(&p1.Id, &p1.Productname, &p1.Image, &p1.Priceold, &p1.Price, &p1.Connectivity)
		if err != nil {
			log.Fatal(err)
		}
		p2 = append(p2, p1)
	}
	return p2

}

func Fetch(id string) detail {

	db, err := sql.Open("mysql", "root:ayush@123@tcp(127.0.0.1:3306)/product")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	results, err := db.Query("select * from productdetails ")
	if err != nil {
		log.Fatal(err)
	}
	var p1 detail
	for results.Next() {

		err = results.Scan(&p1.Id, &p1.Productname, &p1.Image, &p1.Priceold, &p1.Price, &p1.Connectivity)
		if err != nil {
			log.Fatal(err)
		}
		s2 := strconv.Itoa(p1.Id)
		if s2 == id {
			return p1
		}

	}
	return p1
}
