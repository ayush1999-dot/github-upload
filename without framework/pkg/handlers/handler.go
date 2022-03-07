package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	utilities "github.com/ayush1999-dot/go-project/pkg/database"
)

type detail struct {
	Id           int
	Productname  string
	Image        string
	Priceold     int
	Price        int
	Connectivity int
}

var home string
var tpl *template.Template
var tpl2 *template.Template

//var addtocartslice detail

var P3 []detail

type hotdog int

func init() {

	tpl = template.Must(template.ParseGlob("./../htmlandcss/*.html"))
	tpl2 = template.Must(template.ParseGlob("./../htmlandcss(new)/*.html"))

}

var Noofproducts hotdog

/*
func Product(wr http.ResponseWriter, r *http.Request) {

	Products1 := utilities.Getdata("")
	err := tpl.ExecuteTemplate(wr, "index.html", Products1)
	if err != nil {
		log.Fatal(err)
	}
	Fill(r)

}*/
func Cart(wr http.ResponseWriter, r *http.Request) {

	err := tpl.ExecuteTemplate(wr, "cart.html", P3)

	if err != nil {
		log.Fatal(err)
	}
	id := r.FormValue("remove")
	if id != " " {
		Remove(id)
	}
}
func Homepage(wr http.ResponseWriter, r *http.Request) {

	err := tpl2.ExecuteTemplate(wr, "home.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Productpage(wr http.ResponseWriter, r *http.Request) {
	home = r.FormValue("fromnavbar")

	fmt.Println(home)
	Products1 := utilities.Getdata(home)
	err := tpl2.ExecuteTemplate(wr, "productpage.html", Products1)
	if err != nil {
		log.Fatal(err)
	}
}

/*
func Fill(r *http.Request) {
	id := r.FormValue("add")
	if id != "" {
		addtocartslice = detail(utilities.Fetch(id))
		P3 = append(P3, addtocartslice)
	}

}*/
func Remove(id string) {

	for i, s := range P3 {
		if strconv.Itoa(s.Id) == id {
			P3 = append(P3[:i], P3[i+1:]...)

		}
	}

}
