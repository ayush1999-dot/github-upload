package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type detail struct {
	Productname  string
	Image        string
	Priceold     int
	Price        int
	Connectivity int
}
type user struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

/*var home string
var tpl *template.Template
var tpl2 *template.Template
*/

//var addtocartslice detail
const DNS = "root:ayush@123@tcp(127.0.0.1:3306)/pract?charset=utf8mb4&parseTime=True&loc=Local"

var P3 []detail

/*
func init() {

	tpl = template.Must(template.ParseGlob("./../htmlandcss/*.html"))
	tpl2 = template.Must(template.ParseGlob("./../htmlandcss(new)/*.html"))

}*/
var DB *gorm.DB
var err error

func Migrate() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB.AutoMigrate(&user{})

}
func GetUsers(wr http.ResponseWriter, r *http.Request) {

}
func GetUser(wr http.ResponseWriter, r *http.Request) {

}
func CreateUser(wr http.ResponseWriter, r *http.Request) {
	wr.Header().Set("Content-Type", "application/json")
	var users user
	json.NewDecoder(r.Body).Decode(&users)
	DB.Create(&users)
	json.NewEncoder(wr).Encode(users)

}
func UpdateUser(wr http.ResponseWriter, r *http.Request) {

}
func DeleteUser(wr http.ResponseWriter, r *http.Request) {

}

/*
func Product(wr http.ResponseWriter, r *http.Request) {

	Products1 := utilities.Getdata("")
	err := tpl.ExecuteTemplate(wr, "index.html", Products1)
	if err != nil {
		log.Fatal(err)
	}
	Fill(r)

}
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
	Products1 := utilities.Getdata(home)
	err := tpl2.ExecuteTemplate(wr, "productpage.html", Products1)
	if err != nil {
		log.Fatal(err)
	}
}


func Fill(r *http.Request) {
	id := r.FormValue("add")
	if id != "" {
		addtocartslice = detail(utilities.Fetch(id))
		P3 = append(P3, addtocartslice)
	}

}
func Remove(id string) {

	for i, s := range P3 {
		if strconv.Itoa(s.Id) == id {
			P3 = append(P3[:i], P3[i+1:]...)

		}
	}

}*/
