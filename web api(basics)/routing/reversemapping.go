package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ma() {
	r := mux.NewRouter()

	r.Path("/home/{categories}/{id}").HandlerFunc(hanndleHandler)

	url, _ := r.Get("first").URL("categories", "{categories}", "id", "{id}")
	fmt.Println(url.Path)
	http.ListenAndServe(":8080", r)
}
func hanndleHandler(wr http.ResponseWriter, r *http.Request) {}
