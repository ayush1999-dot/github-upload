package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type user struct {
	gorm.Model
	Name      string
	Phoneno   int
	CompanyID int
	Company   Company
}
type Company struct {
	ID   int
	Name string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Company{}, &user{})

	get(db)
}
func get(db *gorm.DB) {
	var user1 user
	var company1 []Company
	// db.first is used to get a single elements by thier id or any firlds written
	db.First(&user1, 2)
	fmt.Println(user1)

	//db.find can be  used to get whole set of data or  a single element
	db.Find(&company1, 2)
	fmt.Println(company1)
	var com1 []user
	//deletes the record which has 15 at its id
	db.Delete(&user{}, 15)
	db.Find(&com1)
	fmt.Println(com1)

}

const dns = "root:ayush@123@tcp(127.0.0.1:3306)/pract?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	check(err)

	Migrate(db)
}
