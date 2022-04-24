package handlers

import (
	db "book_store/pkg/dbcon"
	model "book_store/pkg/models"
	res "book_store/pkg/status"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	Db = db.Connect()

}

func GetAllBook(ctx *gin.Context) {
	book := []model.Book{}
	if err := Db.Select(&book, "SELECT * FROM books where is_deleted != 1 "); err != nil {
		var resi = res.CheckERRor(http.StatusInternalServerError, err)
		ctx.JSON(http.StatusInternalServerError, resi)
		return
	}
	if len(book) == 0 {
		var resi = res.CheckERRor(http.StatusNotFound, sql.ErrNoRows)
		ctx.JSON(http.StatusNotFound, resi)
		return
	}

	resi := res.CheckERRor(http.StatusOK, book)
	ctx.JSON(http.StatusOK, resi)

}

/*---------search books---------------*/
func Searchbook(ctx *gin.Context) {
	book := model.Book{}
	stid := ctx.Param("id")
	id, _ := strconv.Atoi(stid)
	if err := Db.Get(&book, "select * from books where id= ? and  is_deleted != 1;", id); err != nil {
		var resi = res.CheckERRor(http.StatusInternalServerError, err)
		ctx.JSON(http.StatusInternalServerError, resi)
		return
	}
	resi := res.CheckERRor(http.StatusOK, book)
	ctx.JSON(http.StatusOK, resi)

}

/*---------BookByauthor---------------*/
func BookbyAuthor(ctx *gin.Context) {
	author := model.Author{}
	book := []model.Book{}
	firstname := ctx.Query("author_name")
	ctx.String(http.StatusOK, "Hello %s %s", firstname)
	//id, _ := strconv.Atoi(stid)
	if err := Db.Get(&author, "select * from author where author_name = ?; ", fmt.Sprintln(firstname)); err != nil {
		var resi = res.CheckERRor(http.StatusInternalServerError, err)
		ctx.JSON(http.StatusInternalServerError, resi)
		return
	}

	if err := Db.Select(&book, "select *from books where author_id =? and is_deleted!=1;", author.Author_Id); err != nil {
		var resi = res.CheckERRor(http.StatusInternalServerError, err)
		ctx.JSON(http.StatusInternalServerError, resi)
		return
	}

	for _, j := range book {
		ctx.JSON(200, j.Book_name)
	}

}

/* func BookbyPublisher(ctx *gin.Context) {

} */

/*----------update--------------*/
func UpdateBook(ctx *gin.Context) {
	var book model.Book
	ctx.ShouldBindJSON(&book)
	stid := ctx.Param("id")
	id, _ := strconv.Atoi(stid)
	//Db.NamedExec("update books set book_name=? ,book_url=? where id =?")
	rows, err := Db.Prepare("update books set book_name=? ,book_url=? where id =?;")
	if err != nil {
		if err == sql.ErrNoRows {
			var resi = res.CheckERRor(http.StatusNotFound, sql.ErrNoRows)
			ctx.JSON(http.StatusNotFound, resi)
			return
		}
		var resi = res.CheckERRor(http.StatusInternalServerError, err)
		ctx.JSON(http.StatusInternalServerError, resi)
		return

	}
	if _, err := rows.Exec(book.Book_name, book.Book_url, id); err != nil {
		var resi = res.CheckERRor(http.StatusInternalServerError, err)
		ctx.JSON(http.StatusInternalServerError, resi)
		return
	}
}

/*---------display---------------*/
func Display(ctx *gin.Context) {
	book := []model.Book{}
	Db.Select(&book, "SELECT * FROM books where is_deleted != 1 ")
	for _, j := range book {
		ctx.JSON(200, j.ID)
		ctx.JSON(200, j.Book_name)
		c.String

	}

}

/*---------Purchase---------------*/
func PurchaseBook(ctx *gin.Context) {
	//book := model.Book{}
	//stid := ctx.Param("id")
}

/*---------deleteBook---------------
func DeleteBook(ctx *gin.Context) {
	stid := ctx.Param("id")
	id, _ := strconv.Atoi(stid)
	Db.NamedExec("UPDATE  books SET is_deleted = 1 books where id= ?", id)
	Db.NamedExec("UPDATE  books SET is_deleted = 1 books where id= ?", id)
}*/
